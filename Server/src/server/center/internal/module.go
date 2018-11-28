package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	"server/base"
	"server/conf"
	"shared"
	"sync"
)

//运行时角色
type Role struct {
	data  shared.RoleData
	agent gate.Agent
}

type RunMap struct {
	m        *Map
	closeSig chan bool
}

var (
	skeleton  = base.NewSkeleton()
	ChanRPC   = skeleton.ChanRPCServer
	roleInfos = make(map[string]*Role)
	maps      = make(map[int]*RunMap)
)

type Module struct {
	*module.Skeleton
	wg sync.WaitGroup
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton

	//初始化地图
	for i := range conf.Server.MapLoad {
		mapId := conf.Server.MapLoad[i]

		newMap := &RunMap{
			m:        NewMap(mapId),
			closeSig: make(chan bool, 0),
		}

		m.wg.Add(1)
		maps[mapId] = newMap

		go func() {
			log.Debug("map %v goroutine start", mapId)
			newMap.m.Run(newMap.closeSig)
			log.Debug("map %v goroutine exit", mapId)
			m.wg.Done()
		}()
	}
}

func (m *Module) OnDestroy() {
	log.Debug("center module destroy")
	for _, v := range maps {
		v.closeSig <- true
	}
	m.wg.Wait()
}
