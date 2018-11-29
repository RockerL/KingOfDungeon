package internal

import (
	"github.com/name5566/leaf/db/mongodb"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	"github.com/name5566/leaf/network"
	"net"
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

var (
	skeleton                            = base.NewSkeleton()
	ChanRPC                             = skeleton.ChanRPCServer
	roleInfos                           = make(map[string]*Role)
	maps                                = make(map[int32]*Map)
	loadedMapCount                      = 0
	dbSession      *mongodb.DialContext = nil //数据库连接
	ctAgent        *TCPAgent
)

type Module struct {
	*module.Skeleton
	wg sync.WaitGroup
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton

	sessionNum := len(conf.Server.MapLoad) * 2
	//初始化数据连接
	dbSession, err := mongodb.Dial(conf.Server.DBAddr, sessionNum)
	if dbSession == nil {
		log.Error("can not connect mongodb ip %v err %v", conf.Server.DBAddr, err.Error())
		return
	} else {
		log.Release("connect mongodb %v success", conf.Server.DBAddr)
	}

	//连接center server，初始化agent
	conn, err := net.Dial("tcp", conf.Server.CTAddr)
	if err != nil {
		panic(err)
	} else {
		msgParser := network.NewMsgParser()
		msgParser.SetMsgLen(conf.LenMsgLen, 0, 0)
		msgParser.SetByteOrder(conf.LittleEndian)

		ctAgent = NewTCPAgent(conn, msgParser, skeleton.ChanRPCServer)

		go func() {
			ctAgent.Run()

			ctAgent.Close()
			ctAgent.OnClose()
		}()
	}

	//初始化地图和协程，每张地图一个协程
	loadedMapCount = 0
	for i := range conf.Server.MapLoad {
		mapId := conf.Server.MapLoad[i]

		newMap := NewMap(mapId)

		m.wg.Add(1)
		maps[mapId] = newMap

		go func() {
			log.Debug("map %v goroutine start", mapId)
			newMap.Run(newMap.closeSig)
			log.Debug("map %v goroutine exit", mapId)
			m.wg.Done()
		}()

		newMap.Skeleton.ChanRPCServer.Go("LoadMap", newMap)
	}

}

func (m *Module) OnDestroy() {
	log.Debug("center module destroy")
	for _, v := range maps {
		v.closeSig <- true
	}
	m.wg.Wait()
}
