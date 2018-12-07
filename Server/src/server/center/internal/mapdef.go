package internal

import (
	"shared"
	"shared/algorithm"
)

const WorldDBName = "world"       //世界数据库名
const MapTableNamePrefix = "map_" //地图块数据集合表名前缀

//块类型定义
const (
	BlockAir     = 0  //空气
	BlockEarth   = 1  //土块
	BlockWater   = 2  //水块
	BlockGrass   = 3  //草块，顶部是草面，底部是土
	BlockIron    = 4  //铁块
	BlockBronze  = 5  //铜块
	BlockGold    = 6  //金块
	BlockSilver  = 7  //银块
	BlockSulphur = 8  //硫磺块
	BlockCoal    = 9  //煤块
	BlockStone   = 10 //石块
)

const (
	OpDig = 1 //挖操作
)

//地面可交互物体，包含掉落的道具，种植的农作物，野果
type MapObjectData struct {
	ObjectType uint16            `bson:"Type"`
	Pos        algorithm.Vector3 `bson:"Pos"`
	ItemData   shared.RoleItemData
}

//地图块数据定义
type BlockData struct {
	BlockType uint16 `bson:"Type"`
	Content   uint32 `bson:"Durable"`
}

//地图区块数据定义
type MapChunkData struct {
	ChunkId     string                            `bson:"Id"`
	ChunkX      int32                             `bson:"ChunkX"`
	ChunkZ      int32                             `bson:"ChunkZ"`
	BlockArray  [shared.ChunkBlockTotal]BlockData `bson:"BlockArray"`
	ObjectArray []MapObjectData                   `bson:"BlockArray"`
}
