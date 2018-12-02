package shared

import (
	"gopkg.in/mgo.v2/bson"
	"shared/algorithm"
)

//用户信息
type UserData struct {
	Id      string `bson:"Id"`      //运营商的用户ID
	RoleNum uint32 `bson:"RoleNum"` //角色数量
}

//角色信息
type RoleData struct {
	Id     bson.ObjectId     `bson:"_id"`    //数据库ID
	UserId string            `bson:"UserId"` //角色对应的用户ID
	Name   string            `bson:"Name"`   //角色名字
	Sex    uint32            `bson:"Sex"`    //性别
	Level  uint32            `bson:"Level"`  //等级
	MapId  uint32            `bson:"MapId"`  //所在地图ID
	Pos    algorithm.Vector3 `bson:"Pos"`    //当前位置
}

const DBName = "game"        //保存用户数据和角色数据
const UserTableName = "user" //用户数据表名
const RoleTableName = "role" //角色数据表名
const MaxRoleNum = 8         //一个用户最多创建的角色数量

const BlockSize = 4                                               //块的尺寸
const BlockMaxY = 4                                               //游戏中高度方向上的块数量
const ChunkBlockNum = 4                                           //组成区块中单边小块的数量
const MaxChunkNum = 256                                           //地图中单边的区块数量
const MapTotalChunk = MaxChunkNum * MaxChunkNum                   //地图中所有区块数量
const ChunkBlockTotal = ChunkBlockNum * ChunkBlockNum * BlockMaxY //区块中块的数量
const ChunkSize = ChunkBlockNum * BlockSize                       //区块单边尺寸
const WorldSize = ChunkSize * MaxChunkNum                         //地图的单边尺寸

const ClientChunkNum = 3 //客户端需要的区块单边个数，只能为奇数
const ClientChunkTotal = ClientChunkNum * ClientChunkNum
