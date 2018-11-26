using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using UnityEngine;

//角色行为状态
enum CharState
{
    Idle = 0,
    Move = 1,
    Attack = 2,
    Hurt = 3,
    Jump = 4,
}

/// <summary>
/// 角色类
/// </summary>
class Character : Entity
{
    protected CharState mState;

    protected Animator mAnim;

    protected float mRunSpeed = 1;

    public Vector3 targetMovePos { get; set; }

    public Vector3 pos { get { return mObj.transform.position; } }

    virtual public void OnMove()
    {
        Vector3 dir = targetMovePos - mObj.transform.position;
        float dist = dir.magnitude;
        if (dist > 0.01f)
        {
            dir.Normalize();
            mObj.transform.position = Vector3.Lerp(mObj.transform.position, targetMovePos, mRunSpeed * Time.deltaTime / dist);
            mObj.transform.forward = dir;
        }
    }

    virtual public void OnUpdate() { }
}
