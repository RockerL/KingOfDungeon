using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using LitJson;
using UnityEngine;

class BlockCfgItem
{
    public int blockType = 0;
    public double texCoordX = 0.0;
    public double texCoordY = 0.0;
    public double texCoordSizeX = 0.0;
    public double texCoordSizeY = 0.0;
    public uint defaultDurable = 0;
}

class BlocksConfig : Singleton<BlocksConfig>
{
    const string BLOCK_CONFIG = "Config/Blocks.json";    //块的配置数据

    Dictionary<int, BlockCfgItem> mCfgData = new Dictionary<int, BlockCfgItem>();

    Dictionary<int, Vector2[]> mUVs = new Dictionary<int, Vector2[]>();

    public Material material { get; set; } //方块共享的材质
    
    class ConfigData
    {
        public List<BlockCfgItem> items = null;

        public string matPath = null;
    }

    public void Load()
    {
        TextAsset txt = (TextAsset)ResManager.singleton.LoadAsset(BLOCK_CONFIG, typeof(TextAsset));
        ConfigData obj = JsonMapper.ToObject<ConfigData>(txt.text);

        foreach (var v in obj.items)
        {
            mCfgData.Add(v.blockType, v);
            Vector2[] uvs = new Vector2[4];
            Vector2 texCoord = new Vector2((float)v.texCoordX, (float)v.texCoordY); uvs[0] = texCoord;
            texCoord = new Vector2((float)v.texCoordX, (float)v.texCoordY + (float)v.texCoordSizeY); uvs[1] = texCoord;
            texCoord = new Vector2((float)v.texCoordX + (float)v.texCoordSizeX, (float)v.texCoordY + (float)v.texCoordSizeY); uvs[2] = texCoord;
            texCoord = new Vector2((float)v.texCoordX + (float)v.texCoordSizeX, (float)v.texCoordY); uvs[3] = texCoord;
            mUVs.Add(v.blockType, uvs);
        }

        material = (Material)ResManager.singleton.LoadAsset(obj.matPath, typeof(Material));
    }

    public BlockCfgItem GetBlockCfgItem(int type)
    {
        BlockCfgItem item = null;
        if (!mCfgData.TryGetValue(type, out item))
        {
            Debug.LogError("can not find type " + type + " in block config");
            return null;
        }
        return item;
    }

    public Vector2[] GetBlockUVs(int type)
    {
        Vector2[] item = null;
        if (!mUVs.TryGetValue(type, out item))
        {
            Debug.LogError("can not find type " + type + " in block config");
            return null;
        }

        return item;
    }
}

