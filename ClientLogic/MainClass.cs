using System;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using Google.Protobuf;
using System.IO;
using U3DUtility;

public class MainClass
{
    /*
    static private SpObject CreateObject()
    {
        SpObject obj = new SpObject();

        SpObject person = new SpObject();

        SpObject p = new SpObject();
        p.Insert("name", "Alice");
        p.Insert("id", 10000);
        p.Insert("email", "163.com");

        SpObject phone = new SpObject();
        {
            SpObject p1 = new SpObject();
            p1.Insert("number", "123456789");
            p1.Insert("type", 1);
            phone.Append(p1);
        }

        p.Insert("phone", phone);
        person.Append(p);

        obj.Insert("person", person);
        return obj;
    }
    */

    /// <summary>
    /// 逻辑初始化
    /// </summary>
    public static void LogicStart()
    {
        Debug.Log("logic start...");

        BlocksConfig.Instance.Load();
        CameraConfig.Instance.Load();
        PlayerConfig.Instance.Load();

        TcpLayer.singleton.Init(1024 * 1024, 1024 * 1024);

        TcpLayer.singleton.Connect("127.0.0.1", 3566, (isSuccess, msg) =>
        {
            Debug.Log("connect to center " + isSuccess);
        }, (msg)=> {
            Debug.Log("disconneted to server " + msg);
        }, (id, data)=> {
            if (id == 0)
            {
                Msg.role_info info = Msg.role_info.Parser.ParseFrom(data);
                Debug.LogFormat("recv data {0} {1} {2}", id, info.Name, info.Level);

                
            }
        });

        var o = GameObject.Instantiate<GameObject>(Resources.Load("Button") as GameObject);
        o.transform.SetParent(GameObject.Find("Canvas").transform, false);
        Button btn = o.GetComponent<Button>();
        btn.onClick.AddListener( () =>
        {
            /*
            string client_proto = @"
                .Person {
                    name 0 : string
                    id 1 : integer
                    email 2 : string
                    .PhoneNumber {
                        number 0 : string
                        type 1 : integer
                    }
                    phone 3 : *PhoneNumber
                }
                .AddressBook{
                    person 0 : *Person
                }
            ";


            SpTypeManager tm = SpTypeManager.Import(client_proto);

            SpObject obj = CreateObject();

            SpStream encode_stream = tm.Codec.Encode("AddressBook", obj);

            encode_stream.Position = 0;
            SpObject obj2 = tm.Codec.Decode("AddressBook", encode_stream);
            SpObject person = obj2["person"];
            Debug.LogFormat("obj2 {0} ", obj2["person"][0]["name"].AsString());
            */

            PlayerAttrib attrib = new PlayerAttrib();
            ClientScene.Instance.Load(new Vector3(100, 8, 100), attrib);
            /*
            Msg.role_info info = new Msg.role_info();
            info.Level = 100;
            info.Name = "king of dungeon";

            var mem = new MemoryStream();
            info.WriteTo(mem);
            mem.Position = 0;

            TcpLayer.singleton.SendPack(0, mem.ToArray());*/
        });
    }

    /// <summary>
    /// Update总入口函数
    /// </summary>
    public static void GameUpdate()
    {
        if (Input.GetMouseButtonDown(0))
        {
            Ray ray = CameraControl.Instance.GetMouseRay(Input.mousePosition);
            ClientScene.Instance.TryBeginMove(ray);
        }

        ClientScene.Instance.OnUpdate();
    }

    /// <summary>
    /// LateUpdate入口
    /// </summary>
    public static void GameLateUpdate()
    {
        float rotate = 0;
        float pitch = 0;
        float zoom = 0;
        if (Input.GetMouseButton(1))
        {
            rotate = Input.GetAxis("Mouse X");
            pitch = Input.GetAxis("Mouse Y");
        }

        zoom = Input.GetAxis("Mouse ScrollWheel");

        CameraControl.Instance.OnFrameMove(rotate, pitch, zoom);
    }

    /// <summary>
    /// FixedUpdate入口
    /// </summary>
    public static void GameFixedUpdate()
    {

    }

    /// <summary>
    /// 游戏退出时调用
    /// </summary>
    public static void GameQuit()
    {
        ClientScene.Instance.OnQuit();
    }
}
