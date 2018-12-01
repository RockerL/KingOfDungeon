package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/module"
	"server/base"
	"server/conf"
	"shared"
	"shared/algorithm"
)

//运行时角色
type MapRole struct {
	m     *Map
	data  *shared.RoleData
	agent gate.Agent
	idx   int
}

//地图定义
type Map struct {
	*module.Skeleton
	Id       uint32    //地图编号
	closeSig chan bool //地图协程退出信号

	roles     []*MapRole                //地图里容纳的最大角色
	roleIndex *algorithm.IndexAllocator //分配器
}

func NewMap(id uint32) *Map {
	m := &Map{
		Id:        id,
		Skeleton:  base.NewSkeleton(),
		closeSig:  make(chan bool, 0),
		roles:     make([]*MapRole, conf.MapRoleMax),
		roleIndex: algorithm.NewIndexAllocator(conf.MapRoleMax),
	}

	m.ChanRPCServer.Register("LoadMap", handleMapLoad)
	m.ChanRPCServer.Register("RoleEnterMap", handleRoleEnterMap)
	return m
}

func (m *Map) OnInit() {

}

func (m *Map) OnDestroy() {

}
