using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using UnityEngine;
using UnityEngine.AI;
using UnityEngine.Profiling;
using System.Threading;

/// <summary>
/// 客户端场景管理，能看到玩家周围的世界
/// </summary>
class ClientScene : Singleton<ClientScene>
{
    private int mStartChunkX;
    private int mStartChunkZ;

    private int mPlayerChunkX;
    private int mPlayerChunkZ;

    private int mPlayerBlockX = 0;
    private int mPlayerBlockZ = 0;
    
    //所有角色的根物体
    private GameObject mPlayerRoot;
    //当前玩家控制的角色
    private Player mMainPlayer;
    //所有玩家集合
    private Dictionary<int, Player> mAllPlayers = new Dictionary<int, Player>();

    //所有地块的根物体
    private GameObject mChunkRoot;
    //当前可见的地形区块
    private TerrianChunk[] mTerrianChunks = new TerrianChunk[WorldDef.CHUNK_ALL_NUM];
    //缓存中的地形区块
    private Dictionary<string, TerrianChunk> mCacheChunks = new Dictionary<string, TerrianChunk>();
    //所有地块的根对象
    public GameObject chunkRoot { get { return mChunkRoot; } }
    //所有玩家的根对象
    public GameObject playerRoot { get { return mPlayerRoot; } }
    //默认的两个地块类型数据
    private BlockData mBoundryBlock= new BlockData();
    private BlockData mAirBlock = new BlockData();
    //导航网格
    private NavMeshSurface mNavSurface;

    //处于加载中的区块名
    private List<TerrianChunk> mLoadingChunks = new List<TerrianChunk>();

    /// <summary>
    /// 初始化各个子对象
    /// </summary>
    public ClientScene()
    {
        mChunkRoot = new GameObject("ChunkRoot");
        mPlayerRoot = new GameObject("PlayerRoot");

        mBoundryBlock.mType = (byte)WorldDef.WorldBlockType.Boundary;
        mAirBlock.mType = (byte)WorldDef.WorldBlockType.Air;

        mNavSurface = mChunkRoot.AddComponent<NavMeshSurface>();
    }

    /// <summary>
    /// 加载场景
    /// </summary>
    /// <param name="playerPos"></param>
    /// <param name="attrib"></param>
    public void Load(Vector3 pos, PlayerAttrib attrib)
    {
        OnPlayerPosChanged(pos, () => 
        {
            if (mMainPlayer == null)
            {
                mMainPlayer = new Player();
                mMainPlayer.Load(pos, attrib, true);
                mAllPlayers.Add(attrib.id, mMainPlayer);
            }
        });
    }

    /// <summary>
    /// 当主控玩家位置改变时的处理
    /// </summary>
    /// <param name="x"></param>
    /// <param name="z"></param>
    public void OnPlayerPosChanged(Vector3 playerPos, Action onLoadComplete = null)
    {
        int blockX = (int)(playerPos.x / WorldDef.BLOCK_SIZE);
        int blockZ = (int)(playerPos.z / WorldDef.BLOCK_SIZE);
        if (blockX == mPlayerBlockX && blockZ == mPlayerBlockZ)
            return;

        mPlayerBlockX = blockX;
        mPlayerBlockZ = blockZ;

        int chunkX = blockX / WorldDef.CHUNK_BLOCK_NUM;
        int chunkZ = blockZ / WorldDef.CHUNK_BLOCK_NUM;

        //检查玩家是否还在老的区块范围内，如果没有，则无需调整区块
        if (chunkX == mPlayerChunkX && chunkZ == mPlayerChunkZ)
            return;

        Debug.LogFormat("player change chunk {0} {1}", chunkX, chunkZ);

        mPlayerChunkX = chunkX;
        mPlayerChunkZ = chunkZ;

        //玩家改变了当前所在的区块，则需要重新整理区块排列，加载新的区块
        //先把当前区块放入缓存
        for (int i = 0; i < mTerrianChunks.Length; i++)
        {
            PushChunkIntoCache(mTerrianChunks[i]);
        }

        mLoadingChunks.Clear();

        mStartChunkX = chunkX - WorldDef.CHUNK_NUM / 2;
        mStartChunkZ = chunkZ - WorldDef.CHUNK_NUM / 2;
        for (int cz = mStartChunkZ; cz < mStartChunkZ + WorldDef.CHUNK_NUM; cz++)
        {
            for (int cx = mStartChunkX; cx < mStartChunkX + WorldDef.CHUNK_NUM; cx++)
            {
                int chunkIdx = (cz - mStartChunkZ) * WorldDef.CHUNK_NUM + cx - mStartChunkX;

                //尝试从缓存中寻找区块
                string chunkName = GetChunkName(cx, cz);
                TerrianChunk chunk = GetChunkInCache(chunkName);
                if (chunk == null)
                {
                    chunk = new TerrianChunk(cx, cz); //缓存也没有区块则新建一个区块

                    mLoadingChunks.Add(chunk);
                }

                //把新区块写入到该位置
                mTerrianChunks[chunkIdx] = chunk;
            }
        }

        Loom.RunAsync(() =>
        {
            ThreadLoadChunk(onLoadComplete);
        });
    }

    /// <summary>
    /// 根据块世界坐标获得块数据
    /// </summary>
    /// <param name="x"></param>
    /// <param name="y"></param>
    /// <param name="z"></param>
    /// <returns></returns>
    public BlockData GetBlock(int x, int y, int z)
    {
        int chunkX = x / WorldDef.CHUNK_BLOCK_NUM;
        int chunkZ = z / WorldDef.CHUNK_BLOCK_NUM;

        if (x < 0 || y < 0 || z < 0 
            || x >= WorldDef.MAX_BLOCK_NUM 
            || z >= WorldDef.MAX_BLOCK_NUM)
        {
            return mBoundryBlock;
        }

        if (y >= WorldDef.BLOCKMAX_Y)
        {
            return mAirBlock;
        }

        if (chunkX < mStartChunkX || chunkZ < mStartChunkZ
            || chunkX >= mStartChunkX + WorldDef.CHUNK_NUM
            || chunkZ >= mStartChunkZ + WorldDef.CHUNK_NUM)
        {
            return mAirBlock;   //x,z方向上没显示出来的部分，作为空气
        }
        
        return mTerrianChunks[(chunkZ - mStartChunkZ) * WorldDef.CHUNK_NUM + chunkX - mStartChunkX].GetBlock(x, y, z);
    }

    public void OnUpdate()
    {
        foreach (var v in mAllPlayers)
        {
            v.Value.OnUpdate();
        }
    }

    public void OnQuit()
    {

    }

    TerrianChunk FindChunk(int chunkX, int chunkZ)
    {
        for (int i = 0; i < mTerrianChunks.Length; i++)
        {
            if (mTerrianChunks[i].isLoaded && 
                mTerrianChunks[i].chunkX == chunkX && 
                mTerrianChunks[i].chunkZ == chunkZ)
            {
                return mTerrianChunks[i];
            }
        }

        return null;
    }

    TerrianChunk GetChunkInCache(string chunkName)
    {
        TerrianChunk chunk = null;
        if (mCacheChunks.TryGetValue(chunkName, out chunk))
        {
            mCacheChunks.Remove(chunkName);
            chunk.OnGetFromCache();
        }

        return chunk;
    }

    void PushChunkIntoCache(TerrianChunk chunk)
    {
        if (chunk == null)
            return;

        string chunkName = GetChunkName(chunk.chunkX, chunk.chunkZ);
        mCacheChunks[chunkName] = chunk;
        chunk.OnPushIntoCache();
    }

    public string GetChunkName(int chunkX, int chunkZ)
    {
        return chunkX + "_" + chunkZ;
    }

    public void TryBeginMove(Ray ray)
    {
        RaycastHit hit;
        if (Physics.Raycast(ray, out hit))
        {
            mMainPlayer.SetMoveDestination(hit.point);
        }
    }

    /// <summary>
    /// 异步加载一个区块
    /// </summary>
    /// <param name="chunk"></param>
    /// <param name="chunkX"></param>
    /// <param name="chunkZ"></param>
    void ThreadLoadChunk(Action onLoadComplete)
    {
        foreach(var chunk in mLoadingChunks)
        {
            chunk.LoadChunk();
        }

        Loom.QueueOnMainThread(() => {

            foreach (var chunk in mLoadingChunks)
            {
                chunk.OnChunkLoaded();
            }

            Loom.RunAsync(() =>
            {
                ThreadBuildMesh(onLoadComplete);
            });
        });
    }

    /// <summary>
    /// 异步创建网格
    /// </summary>
    void ThreadBuildMesh(Action onLoadComplete)
    {
        //再构造面，否则数据不对
        for (int i = 0; i < mTerrianChunks.Length; i++)
        {
            mTerrianChunks[i].RebuildMeshQuads();
        }

        Loom.QueueOnMainThread(() =>
        {
            for (int i = 0; i < mTerrianChunks.Length; i++)
            {
                mTerrianChunks[i].OnMeshQuadsBuilt();
            }

            //重新生成导航网格
            mNavSurface.BuildNavMesh();

            if (onLoadComplete != null)
            {
                onLoadComplete();
            }
        });
    }
}
