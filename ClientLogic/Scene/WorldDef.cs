using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

class WorldDef
{
    public const int BLOCKMAX_Y = 4;            //游戏中高度方向上的块数量

    public const int CHUNK_NUM = 3;             //客户端可以看到的区块单边数量，奇数

    public const int CHUNK_ALL_NUM = CHUNK_NUM * CHUNK_NUM; //客户端显示出来的区块总数

    public const int CHUNK_BLOCK_NUM = 4;       //组成区块中单边小块的数量

    public const int MAX_CHUNK_NUM = 256;       //场景中单边的区块数量

    public const int MAX_BLOCK_NUM = MAX_CHUNK_NUM * CHUNK_BLOCK_NUM;   //整个场景中的单边块数量

    public const float BLOCK_SIZE = 4;          //小块的大小

    public const int MAX_CHUNK_CACHE = 49;      //缓存的上限

    public const int CHUNK_BLOCK_MAX_NUM = CHUNK_BLOCK_NUM * CHUNK_BLOCK_NUM * BLOCKMAX_Y;  //显示中地块的小块数量

    public const int CHUNK_VERTS_NUM = 24 * CHUNK_BLOCK_MAX_NUM; //显现中的地块中的所有顶点数量

    //游戏中方块的种类
    public enum WorldBlockType
    {
        SolidStart = 0,         //下面的枚举属于实心块，且是系统生成的
        Earth = 1,              //土块
        Stone = 2,              //石块
        Sand = 3,               //沙土块
        Marble = 4,             //大理石
        Gold = 5,               //金矿块
        Silver = 6,             //银矿块
        Iron = 7,               //铁矿块
        Bronze = 8,             //铜矿块
        Sulphur = 9,            //硫磺
        Coal = 10,              //煤矿
        Boundary = 11,          //地图边界块，不可毁坏

        EmptyStart = 30,        //下面的块属于空的块
        Air = 31,               //空气
        Lava = 32,              //岩浆
        River = 33,             //水
        PlayerWall = 34,        //玩家修建的墙和台阶
        PlayerRoomWall = 35,    //玩家修建的房间周围的墙，和房间连接在一起有增益
        PlayerRoom = 36,        //玩家修建的房间，外观上只能看见地表和地表上的家具
        PlayerSteps = 37,       //玩家修建的斜坡台阶，连接各个层
        SafeWall = 38,          //安全区边界块
    }

}