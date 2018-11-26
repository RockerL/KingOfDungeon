using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using UnityEngine;
using LitJson;

class PlayerConfigParam
{
    public List<string> malePrefabs = null;

    public List<string> femalePrefabs = null;

    public List<string> weaponsPrefabs = null;

    public List<string> helmatsPrefabs = null;
}

class PlayerConfig : Singleton<PlayerConfig>
{
    const string Player_CONFIG = "Config/Player.json";    //玩家角色配置数据

    public PlayerConfigParam cfg { get; set; }

    public void Load()
    {
        TextAsset txt = (TextAsset)ResManager.singleton.LoadAsset(Player_CONFIG, typeof(TextAsset));
        cfg = JsonMapper.ToObject<PlayerConfigParam>(txt.text);
    }
}