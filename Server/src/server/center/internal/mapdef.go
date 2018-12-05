package internal

import "shared"

//地图块数据定义
type BlockData struct {
	BlockType uint16 `bson:"Type"`
	Content   uint32 `bson:"Durable"`
}

//地图区块数据定义
type MapChunkData struct {
	ChunkId    string                            `bson:"Id"`
	ChunkX     int32                             `bson:"ChunkX"`
	ChunkZ     int32                             `bson:"ChunkZ"`
	BlockArray [shared.ChunkBlockTotal]BlockData `bson:"BlockArray"`
}

const WorldDBName = "world"       //世界数据库名
const MapTableNamePrefix = "map_" //地图块数据集合表名前缀

const (
	SolidStart = 0 //下面的枚举属于实心块，且是系统生成的
	Earth      = 1 //土块
	Stone      = 2 //石块
	Sand       = 3
	Marble     = 4
	Gold       = 5
	Silver     = 6
	Iron       = 7
	Bronze     = 8
	Sulphur    = 9
	Coal       = 10
	Boundary   = 11

	EmptyStart     = 30 //下面的块属于空的块
	Air            = 31 //空气
	Lava           = 32 //岩浆
	River          = 33 //水
	PlayerWall     = 34 //玩家修建的墙和台阶
	PlayerRoomWall = 35 //玩家修建的房间周围的墙，和房间连接在一起有增益
	PlayerRoom     = 36 //玩家修建的房间，外观上只能看见地表和地表上的家具
	PlayerSteps    = 37 //玩家修建的斜坡台阶，连接各个层
	SafeWall       = 38 //安全区边界块
)

const (
	OP_Dig = 1 //挖操作
)
