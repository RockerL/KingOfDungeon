using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using UnityEngine;
using UnityEngine.AI;

//反序列化的地形块数据
class BlockData
{
    public byte mType = 0;      //类型

    public byte mSubType = 0;   //子类型

    public uint mDurable = 0;   //耐久度

    public bool IsSolidBlock()
    {
        return mType < (byte)WorldDef.WorldBlockType.EmptyStart;
    }
}

//地形的区块，地形由多个区块无缝拼接而成
class TerrianChunk
{
    private Mesh mMesh = new Mesh();    //地形的网格

    private MeshRenderer mMeshRenderer;

    private MeshFilter mMeshFilter;

    private MeshCollider mMeshCollider; //碰撞器

    private GameObject mObj;            //地形游戏物体

    private int mBlockX;               //在大地图中的块索引，可以决定该块的地形块的坐标

    private int mBlockZ;

    private int mChunkX;               //在大地图中的区块索引

    private int mChunkZ;

    private Vector3[] mVerts = new Vector3[WorldDef.CHUNK_VERTS_NUM];

    private Vector2[] mUVs = new Vector2[WorldDef.CHUNK_VERTS_NUM];

    private List<int> mIndices = new List<int>(WorldDef.CHUNK_VERTS_NUM);

    public int blockX { get { return mBlockX; } }

    public int blockZ { get { return mBlockZ; } }

    public int chunkX { get { return mChunkX; } }

    public int chunkZ { get { return mChunkZ; } }

    public bool isLoaded {  get { return mObj != null; } }

    //是否需要重新计算三角面
    public bool isNeedBuildMeshQuad { get; set; }

    //当前加载的小块数据
    private BlockData[] mBlocks = new BlockData[WorldDef.CHUNK_BLOCK_MAX_NUM];

    public TerrianChunk(int chunkX, int chunkZ)
    {
        mChunkX = chunkX;
        mChunkZ = chunkZ;
        mBlockX = chunkX * WorldDef.CHUNK_BLOCK_NUM;
        mBlockZ = chunkZ * WorldDef.CHUNK_BLOCK_NUM;

        for (int i = 0; i < mBlocks.Length; i++)
        {
            mBlocks[i] = new BlockData();
        }

        isNeedBuildMeshQuad = true;
    }

    public void OnChunkLoaded()
    {
        mMesh.vertices = mVerts;

        //生成GameObject
        if (mObj == null)
        {
            mObj = new GameObject(ClientScene.Instance.GetChunkName(mChunkX, mChunkZ));
            mObj.transform.SetParent(ClientScene.Instance.chunkRoot.transform, true);
            mObj.transform.position = new Vector3(mBlockX * WorldDef.BLOCK_SIZE, 0, mBlockZ * WorldDef.BLOCK_SIZE);
        }
    }

    /// <summary>
    /// 加载地块
    /// </summary>
    /// <param name="chunkX"></param>
    /// <param name="chunkZ"></param>
    public void LoadChunk()
    {
        //默认填充数据，一半地面和空气，同时构造顶点数组和纹理坐标数组
        int vertIdx = 0;
        Vector2 c0_0 = new Vector2(0, 0);
        Vector2 c0_1 = new Vector2(0, 1);
        Vector2 c1_1 = new Vector2(1, 1);
        Vector2 c1_0 = new Vector2(1, 0);
        for (int y = 0; y < WorldDef.BLOCKMAX_Y; y++)
        {
            for (int z = 0; z < WorldDef.CHUNK_BLOCK_NUM; z++)
            {
                for (int x = 0; x < WorldDef.CHUNK_BLOCK_NUM; x++)
                {
                    byte type = 0;
                    if (y < WorldDef.BLOCKMAX_Y / 2)
                    {
                        type = (byte)WorldDef.WorldBlockType.Earth;
                    }
                    else
                    {
                        type = (byte)WorldDef.WorldBlockType.Air;
                    }

                    int idx = y * WorldDef.CHUNK_BLOCK_NUM * WorldDef.CHUNK_BLOCK_NUM + z * WorldDef.CHUNK_BLOCK_NUM + x;
                    mBlocks[idx].mType = type;

                    Vector3 start = new Vector3(
                        x * WorldDef.BLOCK_SIZE, 
                        y * WorldDef.BLOCK_SIZE, 
                        z * WorldDef.BLOCK_SIZE);

                    //上下左右前后顺序把顶点加入数组，注意各个面的法线朝外，顶点顺序为按照法线方向的顺时针
                    //上面
                    Vector3 v = new Vector3(0, WorldDef.BLOCK_SIZE, 0) + start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c0_0; vertIdx++;
                    v = new Vector3(0, WorldDef.BLOCK_SIZE, WorldDef.BLOCK_SIZE) + start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c0_1; vertIdx++; 
                    v = new Vector3(WorldDef.BLOCK_SIZE, WorldDef.BLOCK_SIZE, WorldDef.BLOCK_SIZE) + start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c1_1; vertIdx++; 
                    v = new Vector3(WorldDef.BLOCK_SIZE, WorldDef.BLOCK_SIZE, 0) + start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c1_0; vertIdx++;
                    //下面
                    v = start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c0_0; vertIdx++;
                    v = new Vector3(WorldDef.BLOCK_SIZE, 0, 0) + start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c1_0; vertIdx++;
                    v = new Vector3(WorldDef.BLOCK_SIZE, 0, WorldDef.BLOCK_SIZE) + start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c1_1; vertIdx++;
                    v = new Vector3(0, 0, WorldDef.BLOCK_SIZE) + start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c0_1; vertIdx++;
                    //左边
                    v = start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c0_0; vertIdx++;
                    v = new Vector3(0, 0, WorldDef.BLOCK_SIZE) + start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c0_1; vertIdx++;
                    v = new Vector3(0, WorldDef.BLOCK_SIZE, WorldDef.BLOCK_SIZE) + start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c1_1; vertIdx++;
                    v = new Vector3(0, WorldDef.BLOCK_SIZE, 0) + start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c1_0; vertIdx++;
                    //右边
                    v = new Vector3(WorldDef.BLOCK_SIZE, 0, 0) + start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c0_0; vertIdx++; 
                    v = new Vector3(WorldDef.BLOCK_SIZE, WorldDef.BLOCK_SIZE, 0) + start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c0_1; vertIdx++;
                    v = new Vector3(WorldDef.BLOCK_SIZE, WorldDef.BLOCK_SIZE, WorldDef.BLOCK_SIZE) + start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c1_1; vertIdx++;
                    v = new Vector3(WorldDef.BLOCK_SIZE, 0, WorldDef.BLOCK_SIZE) + start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c1_0; vertIdx++;
                    //前面
                    v = new Vector3(0, 0, WorldDef.BLOCK_SIZE) + start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c0_0; vertIdx++;
                    v = new Vector3(WorldDef.BLOCK_SIZE, 0, WorldDef.BLOCK_SIZE) + start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c0_1; vertIdx++; 
                    v = new Vector3(WorldDef.BLOCK_SIZE, WorldDef.BLOCK_SIZE, WorldDef.BLOCK_SIZE) + start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c1_1; vertIdx++; 
                    v = new Vector3(0, WorldDef.BLOCK_SIZE, WorldDef.BLOCK_SIZE) + start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c1_0; vertIdx++;
                    //后面
                    v = start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c0_0; vertIdx++;
                    v = new Vector3(0, WorldDef.BLOCK_SIZE, 0) + start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c0_1; vertIdx++;
                    v = new Vector3(WorldDef.BLOCK_SIZE, WorldDef.BLOCK_SIZE, 0) + start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c1_1; vertIdx++; 
                    v = new Vector3(WorldDef.BLOCK_SIZE, 0, 0) + start;
                    mVerts[vertIdx] = v; mUVs[vertIdx] = c1_0; vertIdx++;
                }
            }
        }
    }

    public void OnMeshQuadsBuilt()
    {
        mMesh.uv = mUVs;
        mMesh.SetIndices(mIndices.ToArray(), MeshTopology.Quads, 0);
        mMesh.RecalculateNormals(); //这里需要重新计算法线

        if (mMeshFilter == null)
        {
            mMeshFilter = mObj.AddComponent<MeshFilter>();
            mMeshFilter.mesh = mMesh;
        }

        //挂接MeshRender
        if (mMeshRenderer == null)
        {
            mMeshRenderer = mObj.AddComponent<MeshRenderer>();
            mMeshRenderer.material = BlocksConfig.Instance.material;
        }

        if (mMeshCollider == null)
        {
            mMeshCollider = mObj.AddComponent<MeshCollider>();
        }

        isNeedBuildMeshQuad = false;
    }

    /// <summary>
    /// 根据数据来构造地形块网格
    /// </summary>
    public void RebuildMeshQuads()
    {
        if (!isNeedBuildMeshQuad)
            return;

        //根据数据构造面：A分为两类block，A类是全实心的block，B类是空的block，当
        //A类和B类邻接时，需要构造平面，y坐标处于最大值及以上，算Air，y坐标小于0的，算A类block
        //x,z坐标超出地图最大范围的，算A类block
        mIndices.Clear();

        Vector2[] uvs_boundary = BlocksConfig.Instance.GetBlockUVs((byte)WorldDef.WorldBlockType.Boundary);

        for (int y = 0; y < WorldDef.BLOCKMAX_Y; y++)
        {
            for (int z = 0; z < WorldDef.CHUNK_BLOCK_NUM; z++)
            {
                for (int x = 0; x < WorldDef.CHUNK_BLOCK_NUM; x++)
                {
                    int idx = y * WorldDef.CHUNK_BLOCK_NUM * WorldDef.CHUNK_BLOCK_NUM + z * WorldDef.CHUNK_BLOCK_NUM + x;
                    int idxStart = (6 * 4) * (y * (WorldDef.CHUNK_BLOCK_NUM * WorldDef.CHUNK_BLOCK_NUM) + z * WorldDef.CHUNK_BLOCK_NUM + x);
                    bool isSelfSolid = mBlocks[idx].IsSolidBlock();

                    if (isSelfSolid)
                    {
                        Vector2[] uvs = BlocksConfig.Instance.GetBlockUVs(mBlocks[idx].mType);

                        bool isTopSolid = ClientScene.Instance.GetBlock(x + mBlockX, y + 1, z + mBlockZ).IsSolidBlock();
                        if (!isTopSolid)
                        {
                            for (int i = idxStart; i < idxStart + 4; i++)
                            { 
                                mIndices.Add(i); mUVs[i] = uvs[i - idxStart];
                            }
                        }
                        bool isDownSolid = ClientScene.Instance.GetBlock(x + mBlockX, y - 1, z + mBlockZ).IsSolidBlock();
                        if (!isDownSolid)
                        {
                            for (int i = idxStart + 4; i < idxStart + 8; i++)
                            {
                                mIndices.Add(i); mUVs[i] = uvs[i - idxStart - 4];
                            }   
                        }
                        bool isLeftSolid = ClientScene.Instance.GetBlock(x + mBlockX - 1, y, z + mBlockZ).IsSolidBlock();
                        if (!isLeftSolid)
                        {
                            for (int i = idxStart + 8; i < idxStart + 12; i++)
                            {
                                mIndices.Add(i); mUVs[i] = uvs[i - idxStart - 8];
                            }
                        }
                        bool isRightSolid = ClientScene.Instance.GetBlock(x + mBlockX + 1, y, z + mBlockZ).IsSolidBlock();
                        if (!isRightSolid)
                        {
                            for (int i = idxStart + 12; i < idxStart + 16; i++)
                            {
                                mIndices.Add(i); mUVs[i] = uvs[i - idxStart - 12];
                            }
                        }
                        bool isFrontSolid = ClientScene.Instance.GetBlock(x + mBlockX, y, z + mBlockZ + 1).IsSolidBlock();
                        if (!isFrontSolid)
                        {
                            for (int i = idxStart + 16; i < idxStart + 20; i++)
                            {
                                mIndices.Add(i); mUVs[i] = uvs[i - idxStart - 16];
                            }
                        }
                        bool isBackSolid = ClientScene.Instance.GetBlock(x + mBlockX, y, z + mBlockZ - 1).IsSolidBlock();
                        if (!isBackSolid)
                        {
                            for (int i = idxStart + 20; i < idxStart + 24; i++)
                            {
                                mIndices.Add(i); mUVs[i] = uvs[i - idxStart - 20];
                            }
                        }
                    }
                    else
                    {
                        byte t = ClientScene.Instance.GetBlock(x + mBlockX, y - 1, z + mBlockZ).mType; //down
                        if (t == (byte)WorldDef.WorldBlockType.Boundary)
                        {
                            for (int i = idxStart + 7; i >= idxStart + 4; i--)
                            {
                                mIndices.Add(i); mUVs[i] = uvs_boundary[i - idxStart - 4];
                            }
                        }

                        t = ClientScene.Instance.GetBlock(x + mBlockX - 1, y, z + mBlockZ).mType; //left
                        if (t == (byte)WorldDef.WorldBlockType.Boundary)
                        {
                            for (int i = idxStart + 11; i >= idxStart + 8; i--)
                            {
                                mIndices.Add(i); mUVs[i] = uvs_boundary[i - idxStart - 8];
                            }
                        }

                        t = ClientScene.Instance.GetBlock(x + mBlockX + 1, y, z + mBlockZ).mType; //right
                        if (t == (byte)WorldDef.WorldBlockType.Boundary)
                        {
                            for (int i = idxStart + 15; i >= idxStart + 12; i--)
                            {
                                mIndices.Add(i); mUVs[i] = uvs_boundary[i - idxStart - 12];
                            }
                        }

                        t = ClientScene.Instance.GetBlock(x + mBlockX, y, z + mBlockZ + 1).mType; //front
                        if (t == (byte)WorldDef.WorldBlockType.Boundary)
                        {
                            for (int i = idxStart + 19; i >= idxStart + 16; i--)
                            {
                                mIndices.Add(i); mUVs[i] = uvs_boundary[i - idxStart - 16];
                            }
                        }

                        t = ClientScene.Instance.GetBlock(x + mBlockX, y, z + mBlockZ - 1).mType; //back
                        if (t == (byte)WorldDef.WorldBlockType.Boundary)
                        {
                            for (int i = idxStart + 23; i >= idxStart + 20; i--)
                            {
                                mIndices.Add(i); mUVs[i] = uvs_boundary[i - idxStart - 20];
                            }
                        }
                    }
                }
            }
        }
    }

    public BlockData GetBlock(int x, int y, int z)
    {
        if (x < mBlockX || z < mBlockZ || y < 0 
            || x >= mBlockX + WorldDef.CHUNK_BLOCK_NUM
            || z >= mBlockZ + WorldDef.CHUNK_BLOCK_NUM
            || y >= WorldDef.BLOCKMAX_Y)
        {
            Debug.LogFormat("Chunk GetBlock param invalid: {0} {1} {2} start {3} {4}", x, y, z, mBlockX, mBlockZ);
            return null;
        }

        int idx = y * WorldDef.CHUNK_BLOCK_NUM * WorldDef.CHUNK_BLOCK_NUM + (z - mBlockZ) * WorldDef.CHUNK_BLOCK_NUM + (x - mBlockX);
        return mBlocks[idx];
    }

    public void OnPushIntoCache()
    {
        if (mObj != null)
            mObj.SetActive(false);
    }

    public void OnGetFromCache()
    {
        if (mObj != null)
            mObj.SetActive(true);
    }
}
