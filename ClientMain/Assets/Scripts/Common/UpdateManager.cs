using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using UnityEngine;
using System.IO;

public class UpdateManager : U3DSingleton<UpdateManager>
{
    enum UpdateStage
    {
        CheckDownloadIndex, //下载并检查索引文件，生成下载列表
        Downloading,        //下载需要更新的资源包
        LoadLuaScript,      //加载Lua资源包
    }

    public delegate void ProcessCompleteEvent();

    private UpdateStage mStage = UpdateStage.CheckDownloadIndex;
    private string mHttpAddress;
    private List<BundleItem> mDownloadingList = new List<BundleItem>();
    private int mTotalDownloadBytes = 0;
    private int mCurrentDownloadIdx = 0;
    private int mAlreadyDownloadBytes = 0;
    private WWW mWWW = null;
    private string mNewIndexContent;
    private ProcessCompleteEvent mAllDoneEvent;

    /// <summary>
    /// 获取下载进度
    /// </summary>
    public float downloadingProgress
    {
        get
        {
            int currentBytes = 0;
            if (mWWW != null && mCurrentDownloadIdx < mDownloadingList.Count)
            {
                currentBytes = (int)(mDownloadingList[mCurrentDownloadIdx].m_FileSize * mWWW.progress);
            }

            if (mTotalDownloadBytes > 0)
            {
                return (float)(mAlreadyDownloadBytes + currentBytes) / (float)mTotalDownloadBytes;
            }

            return 0;
        }
    }

    /// <summary>
    /// 获取热更新进度
    /// </summary>
    public float totalProgress
    {
        get
        {
            if (mStage == UpdateStage.CheckDownloadIndex)
                return 0;
            else if (mStage == UpdateStage.Downloading)
                return 0.1f + downloadingProgress * 0.8f;
            else if (mStage == UpdateStage.LoadLuaScript)
                return 0.9f;
            else
                return 1;
        }
    }

    /// <summary>
    /// 开启热更新
    /// </summary>
    /// <param name="httpServerIP"></param>
    public void StartUpdate(string httpServerIP, ProcessCompleteEvent allDoneEv)
    {
        Debug.Log("start update resource from " + httpServerIP);

        mHttpAddress = "http://" + httpServerIP + "/" + ResUtils.BundleRootDirName + '/';
        mAllDoneEvent = allDoneEv;
        mStage = UpdateStage.CheckDownloadIndex;

        StartCoroutine(AsyncCheckDownloadingList(OnCompleteCheckDownloadList));
    }

    void OnCompleteCheckDownloadList()
    {
        mStage = UpdateStage.Downloading;

        StartCoroutine(AsyncDownloading(OnCompleteDownloading));
    }

    void OnCompleteDownloading()
    {
        mStage = UpdateStage.LoadLuaScript;

        StartCoroutine(AsyncLoadDLL(OnCompleteLoadLua));
    }

    void OnCompleteLoadLua()
    {
        Debug.Log("update resource complete...");

        if (mAllDoneEvent != null)
            mAllDoneEvent.Invoke();
    }

    /// <summary>
    /// 从服务器得到资源列表并对比出需要更新的包列表
    /// </summary>
    /// <param name="ev">检查完成后回调函数</param>
    /// <returns></returns>
    IEnumerator AsyncCheckDownloadingList(ProcessCompleteEvent ev)
    {
        //读取本地的idx和apk里的idx文件
        Dictionary<string, BundleItem> localBundlesDict = new Dictionary<string, BundleItem>();
        string localIndexPath = ResUtils.BundleRootPath + ResUtils.BundleIndexFileName;

        if (!File.Exists(localIndexPath)) //如果P目录里没有索引文件，去Resources里拷贝一份到P目录
        {
            UnityEngine.Debug.Log("local idx not found, try copy from default");
            Directory.CreateDirectory(ResUtils.BundleRootPath);
            var txt = Resources.Load(ResUtils.BundleIndexFileName.Substring(ResUtils.BundleIndexFileName.IndexOf('.'))) as TextAsset;
            if (txt != null)
                File.WriteAllText(ResUtils.BundleRootPath + ResUtils.BundleIndexFileName, txt.text);
        }

        if (File.Exists(localIndexPath))
        {
            string indexContent = File.ReadAllText(localIndexPath);
            if (indexContent != null)
            {
                IdxFile file = new IdxFile();
                List<BundleItem> list = file.Load(indexContent);
                foreach (var v in list)
                {
                    localBundlesDict[v.m_Name] = v;
                }
            }
        }
        else
        {
            UnityEngine.Debug.LogWarning("local idx not found");
        }

        //下载网上的idx文件
        WWW www = new WWW(mHttpAddress + ResUtils.GetBundleManifestName(Application.platform) + "/" + ResUtils.BundleIndexFileName);
        yield return www;

        if (www.error != null)
            UnityEngine.Debug.Log("remote idx read error " + www.error);

        mDownloadingList.Clear();

        if (www.error == null)
        {
            mNewIndexContent = www.text;
            IdxFile file = new IdxFile();
            List<BundleItem> listServer = file.Load(mNewIndexContent);
            foreach (var v in listServer)
            {
                string localHash = null;
                string netHash = v.m_HashCode;
                BundleItem localItem = null;
                if (localBundlesDict.TryGetValue(v.m_Name, out localItem))
                    localHash = localItem.m_HashCode;

                if (localHash != netHash)
                    mDownloadingList.Add(v); //网上的资源较新则需要重新下载到本地
            }

            UnityEngine.Debug.LogFormat("download idx file success! new bundles count {0}, downloading {1}", listServer.Count, mDownloadingList.Count);
        }
        else
        {
            UnityEngine.Debug.LogFormat("download idx file error! {0}", www.error);
        }

        if (ev != null)
            ev.Invoke();

        yield return null;
    }

    /// <summary>
    /// 异步下载需要更新的资源
    /// </summary>
    /// <param name="ev">下载完成回调函数</param>
    /// <returns></returns>
    IEnumerator AsyncDownloading(ProcessCompleteEvent ev)
    {
        mTotalDownloadBytes = 0;
        mCurrentDownloadIdx = 0;
        mAlreadyDownloadBytes = 0;
        foreach (var v in mDownloadingList)
        {
            mTotalDownloadBytes += v.m_FileSize;
        }

        foreach (var v in mDownloadingList)
        {
            string url = mHttpAddress + ResUtils.GetBundleManifestName(Application.platform) + "/" + v.m_Name;
            UnityEngine.Debug.LogFormat("downloading {0} size {1}", v.m_Name, v.m_FileSize);
            WWW www = new WWW(url);
            mWWW = www;
            yield return www;
            if (www.error == null)
            {
                string fileName = ResUtils.BundleRootPath + v.m_Name;
                string dir = fileName.Substring(0, fileName.LastIndexOf('/'));
                Directory.CreateDirectory(dir);
                File.WriteAllBytes(fileName, www.bytes);
            }
            else
            {
                UnityEngine.Debug.LogErrorFormat("downloading {0} error {1}", v.m_Name, www.error);
            }
            mAlreadyDownloadBytes += v.m_FileSize;
            mCurrentDownloadIdx++;
        }

        //全部下载成功后，再覆盖写入索引文件
        Directory.CreateDirectory(ResUtils.BundleRootPath);
        if (mNewIndexContent != null)
        {
            File.WriteAllText(ResUtils.BundleRootPath + ResUtils.BundleIndexFileName, mNewIndexContent);
            mNewIndexContent = null;
        }

        if (ev != null)
            ev.Invoke();

        yield return null;
    }

    /// <summary>
    /// 从bundle中异步加载lua文件
    /// </summary>
    /// <param name="ev">加载完毕后回调</param>
    /// <returns></returns>
    IEnumerator AsyncLoadDLL(ProcessCompleteEvent ev)
    {
        string filePath = ResUtils.BundleRootPath + "ClientLogic.unity3d";

        var fileInfo = new FileInfo(filePath);
        if (fileInfo == null)
            yield break;

        AssetBundle bundle = AssetBundle.LoadFromFile(fileInfo.FullName);
        if (bundle == null)
            yield break;

        AssetBundleRequest request = bundle.LoadAllAssetsAsync();
        yield return request;

        if (request.allAssets.Length == 0)
            yield break;

        var text = request.allAssets[0] as TextAsset;
        if (text == null)
            yield break;

        using (System.IO.MemoryStream fs = new MemoryStream(text.bytes))
        {
            GameMain.domain.LoadAssembly(fs); 
        }

        if (ev != null)
            ev.Invoke();

        yield return null;
    }
   
}

