using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using UnityEngine;

class CameraControl : Singleton<CameraControl>
{
    private Transform mMainPlayer;

    private Camera mCamera;

    private float mCurDeltaPitch = 0;

    public void Init(Transform player)
    {
        mCamera = Camera.main;
        mMainPlayer = player;

        if (mCamera == null || mMainPlayer == null)
        {
            Debug.LogError("main camera not found");
            return;
        }

        mCamera.transform.SetParent(mMainPlayer.transform, false);

        Vector3 pos = mMainPlayer.transform.position;
        pos = pos - (mMainPlayer.transform.forward * (float)CameraConfig.Instance.cfg.initBackDist);
        pos = pos + Vector3.up * (float)CameraConfig.Instance.cfg.initHeight;
        mCamera.transform.position = pos;

        Vector3 dir = mMainPlayer.transform.position - pos;
        dir.Normalize();
        mCamera.transform.forward = dir;

        mCurDeltaPitch = Mathf.Rad2Deg * Mathf.Atan2((float)CameraConfig.Instance.cfg.initHeight, (float)CameraConfig.Instance.cfg.initBackDist);
    }

    public void OnFrameMove(float rotate, float pitch, float zoom)
    {
        if (mMainPlayer == null)
            return;

        OnRotate(rotate);

        OnPitch(pitch);

        OnZoom(zoom);
    }

    public Ray GetMouseRay(Vector3 mousePos)
    {
        if (mCamera == null)
            return new Ray();

        return mCamera.ScreenPointToRay(mousePos);
    }

    void OnRotate(float rotate)
    {
        if (Mathf.Abs(rotate) < 0.2f)
            return;

        Vector3 dir = mMainPlayer.transform.position - mCamera.transform.position;
        float rad = Mathf.Atan2(dir.z, dir.x);
        float radius = Mathf.Sqrt(dir.x * dir.x + dir.z * dir.z);

        rad -= (Time.deltaTime * rotate * (float)CameraConfig.Instance.cfg.rotateSpeed);

        Vector3 newPos = mCamera.transform.position;
        newPos.x = mMainPlayer.transform.position.x - radius * Mathf.Cos(rad);
        newPos.z = mMainPlayer.transform.position.z - radius * Mathf.Sin(rad);
        mCamera.transform.position = newPos;

        dir = mMainPlayer.transform.position - mCamera.transform.position;
        dir.Normalize();
        mCamera.transform.forward = dir;
    }

    void OnPitch(float pitch)
    {
        if (Math.Abs(pitch) < 0.2f)
            return;

        pitch *= (float)CameraConfig.Instance.cfg.pitchSpeed;

        float targetPitch = mCurDeltaPitch - pitch;
        if (targetPitch > CameraConfig.Instance.cfg.pitchMaxDegree || targetPitch < CameraConfig.Instance.cfg.pitchMinDegree)
            return;
        
        mCamera.transform.RotateAround(mMainPlayer.transform.position, -mCamera.transform.right, pitch);
        mCurDeltaPitch = targetPitch;
    }

    void OnZoom(float zoom)
    {
        if (Math.Abs(zoom) < 0.2f)
            return;

        Vector3 dir = mMainPlayer.transform.position - mCamera.transform.position;
        float targetLen = dir.magnitude + zoom;

        if (targetLen > CameraConfig.Instance.cfg.maxDist || targetLen < CameraConfig.Instance.cfg.minDist)
            return;

        dir.Normalize();
        dir *= targetLen;
        mCamera.transform.position = mMainPlayer.transform.position - dir;
    }
}