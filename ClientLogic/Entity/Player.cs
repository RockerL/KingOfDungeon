using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using UnityEngine;
using UnityEngine.AI;

class PlayerAttrib
{
    public int id = -1;

    public int lifePoint = 0;       //生命值
    public byte level = 0;          //等级
    public int curExp = 0;          //经验值
    public byte sex = 0;            //性别
    
    //下面的字段表示角色外观
    public byte baseMod = 0;        //基础模型
    public byte weapon = 0;         //拿的武器 = 0空手
    public byte helm = 0;           //头盔=0 未戴
    public byte face = 0;           //表情=0 默认
    public byte wing = 0;           //翅膀=0 未戴
    public byte bag = 0;            //背包=0 未戴
    public byte suit = 0;           //外套=0 未戴
}

/// <summary>
/// 玩家类
/// </summary>
class Player : Character
{
    PlayerAttrib mAttrib;

    bool mIsControlPlayer = false;

    NavMeshAgent mNavAgent;

    /// <summary>
    /// 加载玩家到场景中
    /// </summary>
    /// <param name="pos"></param>
    /// <param name="attrib"></param>
    /// <param name="isControlPlayer"></param>
    public void Load(Vector3 pos, PlayerAttrib attrib, bool isControlPlayer)
    {
        mAttrib = attrib;
        mIsControlPlayer = isControlPlayer;

        //从配置文件中加载角色模型
        string baseModPath;
        if (mAttrib.sex == 0)
            baseModPath = PlayerConfig.Instance.cfg.malePrefabs[mAttrib.baseMod];
        else
            baseModPath = PlayerConfig.Instance.cfg.femalePrefabs[mAttrib.baseMod];

        GameObject res = (GameObject)ResManager.singleton.LoadAsset(baseModPath, typeof(GameObject));

        mObj = GameObject.Instantiate<GameObject>(res);
        mObj.transform.SetParent(ClientScene.Instance.playerRoot.transform, false);
        mObj.transform.position = pos;

        mNavAgent = mObj.AddComponent<NavMeshAgent>();

        mAnim = mObj.GetComponent<Animator>();

        if (mIsControlPlayer)
        {
            CameraControl.Instance.Init(mObj.transform);
        }
    }

    public void SetMoveDestination(Vector3 pos)
    {
        bool isSet = mNavAgent.SetDestination(pos);
        if(isSet)
        {
            mState = CharState.Move;
            mNavAgent.isStopped = false;
            ChangeStateAnimation();
        }
    }

    public override void OnUpdate()
    {
        if (mState == CharState.Move)
        {
            if (mNavAgent.remainingDistance < 0.1f)
            {
                mNavAgent.isStopped = true;
                mState = CharState.Idle;
                ChangeStateAnimation();
            }

            ClientScene.Instance.OnPlayerPosChanged(mObj.transform.position);
        }
    }

    void ChangeStateAnimation()
    {
        if (mAnim == null)
            return;

        switch(mState)
        {
            case CharState.Idle:
                mAnim.SetBool("Run", false);
                break;
            case CharState.Move:
                mAnim.SetBool("Run", true);
                break;
        }
    }
}