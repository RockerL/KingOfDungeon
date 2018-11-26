using UnityEngine;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using ILRuntime.Runtime.Enviorment;

/// <summary>
/// 主工程运行时入口类
/// </summary>
public class GameMain : MonoBehaviour
{
    //AppDomain是ILRuntime的入口，最好是在一个单例类中保存，整个游戏全局就一个
    static AppDomain mAppDomain;

    static public GameMain instance { get; set; }

    public static AppDomain domain
    {
        get { return mAppDomain;  }
    }

    private void Awake()
    {
        instance = this;
    }

    void Start()
    {
        //首先实例化ILRuntime的AppDomain，AppDomain是一个应用程序域，每个AppDomain都是一个独立的沙盒
        if (mAppDomain == null)
        {
            mAppDomain = new ILRuntime.Runtime.Enviorment.AppDomain();

            mAppDomain.DebugService.StartDebugService(56000);

            StartCoroutine(LoadHotFixAssembly());
        }
    }

    IEnumerator LoadHotFixAssembly()
    {
#if UNITY_EDITOR
        WWW www = new WWW("file:///" + Application.streamingAssetsPath + "/ClientLogic.dll");
#else
        WWW www = new WWW(Application.persistentDataPath + "/ClientLogic.dll");
#endif

        while (!www.isDone)
        {
            yield return null;
        }
            
        if (!string.IsNullOrEmpty(www.error))
        {
            UnityEngine.Debug.LogError(www.error);
        }
            
        byte[] dllBytes = www.bytes;
        www.Dispose();

        //PDB文件是调试数据库，如需要在日志中显示报错的行号，则必须提供PDB文件，不过由于会额外耗用内存，正式发布时请将PDB去掉
#if UNITY_EDITOR
        www = new WWW("file:///" + Application.streamingAssetsPath + "/ClientLogic.pdb");
#else
        www = new WWW(Application.persistentDataPath + "/ClientLogic.pdb");
#endif
        while (!www.isDone)
        { 
            yield return null;
        }

        if (!string.IsNullOrEmpty(www.error))
        { 
            UnityEngine.Debug.LogError(www.error);
        }

        byte[] pdbBytes = www.bytes;

        using (System.IO.MemoryStream fs = new MemoryStream(dllBytes))
        {
            using (System.IO.MemoryStream p = new MemoryStream(pdbBytes))
            {
                mAppDomain.LoadAssembly(fs, p, new Mono.Cecil.Pdb.PdbReaderProvider());
            }
        }

        InitializeILRuntime();

        OnHotFixLoaded();
    }

    void InitializeILRuntime()
    {

        //初始化CLR绑定，让DLL里面的调用更快
        ILRuntime.Runtime.Generated.CLRBindings.Initialize(mAppDomain);
        //注册LitJson到DLL
        LitJson.JsonMapper.RegisterILRuntimeCLRRedirection(mAppDomain);
        //注册一些类到DLL
        ColliderListener.RegisterILRuntime(mAppDomain);
        U3DUtility.TcpLayer.RegisterILRuntime(mAppDomain);

        //注册MonoBehaviour到DLL
        mAppDomain.RegisterCrossBindingAdaptor(new MonoBehaviourAdapter());
        //注册协程到DLL
        mAppDomain.RegisterCrossBindingAdaptor(new CoroutineAdapter());
        //注册Protobuf
        mAppDomain.DelegateManager.RegisterFunctionDelegate<Adapt_IMessage.Adaptor>();
        mAppDomain.RegisterCrossBindingAdaptor(new Adapt_IMessage());
    }

    void OnHotFixLoaded()
    {
        mAppDomain.Invoke("MainClass", "LogicStart", null, null);
    }

    private void Update()
    {
        mAppDomain.Invoke("MainClass", "GameUpdate", null, null);
    }

    private void LateUpdate()
    {
        mAppDomain.Invoke("MainClass", "GameLateUpdate", null, null);
    }

    private void FixedUpdate()
    {
        mAppDomain.Invoke("MainClass", "GameFixedUpdate", null, null);
    }

    private void OnApplicationQuit()
    {
        mAppDomain.Invoke("MainClass", "GameQuit", null, null);
    }
}
