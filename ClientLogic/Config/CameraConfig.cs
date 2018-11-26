using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using UnityEngine;
using LitJson;

class CameraParam
{
    public double initBackDist = 10;        //初始时距离角色身后的距离
    public double initHeight = 10;          //初始时比角色高多少，这两个参数同时也决定了相机的初始角度
    public double minDist = 5;              //距离角色至少的距离
    public double maxDist = 15;             //距离角色最多的距离
    public double rotateSpeed = 3;          //相机旋转速度
    public double pitchSpeed = 3;           //相机俯仰速度
    public double pitchMaxDegree = 60;      //相机俯仰最大相对地面角度
    public double pitchMinDegree = 45;      //相机俯仰最小距离地面角度
}

class CameraConfig : Singleton<CameraConfig>
{
    const string CAMERA_CONFIG = "Config/Camera.json";    //相机配置文件路径

    public CameraParam cfg { get; set; }

    public void Load()
    {
        TextAsset txt = (TextAsset)ResManager.singleton.LoadAsset(CAMERA_CONFIG, typeof(TextAsset));
        cfg = JsonMapper.ToObject<CameraParam>(txt.text);
    }

}