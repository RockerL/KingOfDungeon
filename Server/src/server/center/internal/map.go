package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/module"
	"proto"
	"server/base"
	"server/conf"
	"shared"
	"shared/algorithm"
)

//运行时地图定义
type Map struct {
	*module.Skeleton
	Id       uint32    //地图编号
	closeSig chan bool //地图协程退出信号

	roles     []*MapRole                //地图里容纳的最大角色
	roleIndex *algorithm.IndexAllocator //角色索引分配器

	chunks []MapChunk //地图的所有运行时区块
}

func NewMap(id uint32) *Map {
	m := &Map{
		Id:        id,
		Skeleton:  base.NewSkeleton(),
		closeSig:  make(chan bool, 0),
		roles:     make([]*MapRole, conf.MapRoleMax),
		roleIndex: algorithm.NewIndexAllocator(conf.MapRoleMax),
		chunks:    make([]MapChunk, shared.MaxChunkNum*shared.MaxChunkNum),
	}

	//注册地图协程的消息处理函数
	m.ChanRPCServer.Register("LoadMap", handleMapLoad)
	m.ChanRPCServer.Register("RoleEnterMap", handleRoleEnterMap)
	return m
}

func (m *Map) OnInit() {

}

func (m *Map) OnDestroy() {

}

func (m *Map) GetChunk(chunkX int32, chunkZ int32) *MapChunk {
	if chunkX < 0 || chunkZ < 0 ||
		chunkX >= shared.MaxChunkNum ||
		chunkZ >= shared.MaxChunkNum {
		return nil
	}

	chunkIdx := chunkZ*shared.MaxChunkNum + chunkX

	return &m.chunks[chunkIdx]
}

func (m *Map) RoleEnter(agent gate.Agent) {

	mapRole := agent.UserData().(*MapRole)

	//确定起始同步的chunk索引
	chunkX := int32(mapRole.data.Pos.X / shared.ChunkSize)
	chunkZ := int32(mapRole.data.Pos.Z / shared.ChunkSize)

	startChunkX := chunkX - 1
	startChunkZ := chunkZ - 1

	chunkNum := shared.ClientChunkNum * shared.ClientChunkNum

	rsp := &proto.RspEnterGs{
		RetCode:     0,
		MainRoleIdx: mapRole.idx,
		MainRole:    mapRole.MakeBaseInfo(),
		Chunks:      make([]*proto.ChunkInfo, chunkNum),
	}

	for z := int32(0); z < shared.ClientChunkNum; z++ {
		for x := int32(0); x < shared.ClientChunkNum; x++ {
			idx := z*shared.ClientChunkNum + x
			chunk := m.GetChunk(x+startChunkX, z+startChunkZ)
			if chunk != nil {
				rsp.Chunks[idx] = chunk.MakeChunkInfo()
			}
		}
	}

	agent.WriteMsg(rsp)
}
