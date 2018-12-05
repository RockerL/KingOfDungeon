package internal

import (
	"fmt"
	"github.com/name5566/leaf/db/mongodb"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	"gopkg.in/mgo.v2/bson"
	"proto"
	"server/base"
	"server/conf"
	"shared"
	"shared/algorithm"
	"strconv"
)

type WaitLeaveRole struct {
	role         *MapRole
	remainSecond int32
}

//运行时地图定义
type Map struct {
	*module.Skeleton
	Id               uint32                         //地图编号
	closeSig         chan bool                      //地图协程退出信号
	rolesArray       []*MapRole                     //地图里容纳的最大角色
	roleIdxAllocator *algorithm.IndexAllocator      //角色索引分配器
	rolesMap         map[bson.ObjectId]*MapRole     //所有正在玩的角色查找表
	chunks           [shared.MapTotalChunk]MapChunk //地图的所有运行时区块
	collectionName   string                         //地图的数据库集合名
	dbSession        *mongodb.Session               //数据库会话，一个map协程一个
	waitLeaveRoles   map[string]*WaitLeaveRole      //等待离开地图的角色
}

func NewMap(id uint32) *Map {
	m := &Map{
		Id:               id,
		Skeleton:         base.NewSkeleton(),
		closeSig:         make(chan bool, 0),
		rolesArray:       make([]*MapRole, conf.MapRoleMax),
		rolesMap:         make(map[bson.ObjectId]*MapRole),
		roleIdxAllocator: algorithm.NewIndexAllocator(conf.MapRoleMax),
		collectionName:   MapTableNamePrefix + strconv.Itoa(int(id)),
		dbSession:        DBSession.Ref(),
		waitLeaveRoles:   make(map[string]*WaitLeaveRole),
	}

	//注册地图协程的消息处理函数
	m.ChanRPCServer.Register("LoadMap", mapHandleMapLoad)
	m.ChanRPCServer.Register("RoleEnterMap", mapHandleRoleEnterMap)
	m.ChanRPCServer.Register("RoleDisconnect", mapHandleRoleDisconnect)
	m.ChanRPCServer.Register("RoleAction", mapHandleRoleAction)
	m.ChanRPCServer.Register("RoleOpBlock", mapHandleRoleOpBlock)
	m.ChanRPCServer.Register("OnTimerSecond", mapHandleOnTimerSecond)
	return m
}

func (m *Map) OnInit() {

}

func (m *Map) OnDestroy() {
	DBSession.UnRef(m.dbSession)
}

//获取区块，如果没有初始化则初始化
func (m *Map) GetChunk(chunkX int32, chunkZ int32) *MapChunk {
	if chunkX < 0 || chunkZ < 0 ||
		chunkX >= shared.MaxChunkNum ||
		chunkZ >= shared.MaxChunkNum {
		return nil
	}

	chunkIdx := chunkZ*shared.MaxChunkNum + chunkX
	chunk := &m.chunks[chunkIdx]

	if chunk.data.ChunkId == "" {
		chunkName := fmt.Sprintf("%v_%v", chunkX, chunkZ)
		chunkData := &chunk.data
		err := m.dbSession.DB(WorldDBName).C(m.collectionName).Find(bson.M{"Id": chunkName}).One(&chunkData)
		if err != nil {
			chunk.InitChunkWithFlat(chunkName, chunkX, chunkZ)
			m.dbSession.DB(WorldDBName).C(m.collectionName).Insert(chunkData)
		}
	}

	return chunk
}

//角色进入地图
func (m *Map) RoleEnter(mapRole *MapRole) {
	m.rolesArray[mapRole.idx] = mapRole
	m.rolesMap[mapRole.data.Id] = mapRole

	chunkX := int32(mapRole.data.Pos.X / shared.ChunkSize)
	chunkZ := int32(mapRole.data.Pos.Z / shared.ChunkSize)

	roleChunk := m.GetChunk(chunkX, chunkZ)
	mapRole.ChangeChunk(nil, roleChunk)

	rsp := &proto.RspEnterGs{
		RetCode:     0,
		MainRoleIdx: mapRole.idx,
		MainRole:    mapRole.MakeBaseInfo(),
	}
	if mapRole.agent != nil {
		mapRole.agent.WriteMsg(rsp)
	}
}

//角色离开地图
func (m *Map) RoleLeave(mapRole *MapRole) {
	chunkX := int32(mapRole.data.Pos.X / shared.ChunkSize)
	chunkZ := int32(mapRole.data.Pos.Z / shared.ChunkSize)
	roleChunk := m.GetChunk(chunkX, chunkZ)

	if roleChunk != nil {
		mapRole.ChangeChunk(roleChunk, nil)
	} else {
		log.Error("role leave from nil chunk")
	}

	m.rolesArray[mapRole.idx] = nil
	delete(m.rolesMap, mapRole.data.Id)
	m.roleIdxAllocator.Free(int(mapRole.idx))
}

//广播消息给角色周围的区块
func (m *Map) BroadcastAroundRole(chunk *MapChunk, msg interface{}) {
	chunkStartX := chunk.data.ChunkX - int32(shared.ClientChunkNum)/2
	chunkStartZ := chunk.data.ChunkZ - int32(shared.ClientChunkNum)/2

	for z := chunkStartZ; z < chunkStartZ+shared.ClientChunkNum; z++ {
		for x := chunkStartX; x < chunkStartX+shared.ClientChunkNum; x++ {
			chunk := m.GetChunk(x, z)
			if chunk != nil {
				chunk.roles.Traversal(func(v algorithm.ElemType) {
					r := v.(*MapRole)
					if r.agent != nil {
						r.agent.WriteMsg(msg)
					}
				})
			}
		}
	}
}
