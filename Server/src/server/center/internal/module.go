package internal

import (
	"github.com/name5566/leaf/db/mongodb"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	"github.com/name5566/leaf/network"
	"net"
	"proto"
	"server/base"
	"server/conf"
	"shared"
	"sync"
)

var (
	skeleton                            = base.NewSkeleton()
	ChanRPC                             = skeleton.ChanRPCServer
	waitEnterRoles                      = make(map[string]int64)
	Maps                                = make(map[uint32]*Map)
	loadedMapCount                      = 0
	DBSession      *mongodb.DialContext = nil //数据库连接
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
		DBSession = dbSession
	}

	//设置center server to game server 消息路由
	shared.GSCTProcessor.SetRouter(&proto.NotifyRoleEnter{}, ChanRPC)

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
		Maps[mapId] = newMap

		go func() {
			log.Debug("map %v goroutine start", mapId)
			newMap.Run(newMap.closeSig)
			log.Debug("map %v goroutine exit", mapId)
			m.wg.Done()
		}()

		newMap.ChanRPCServer.Go("LoadMap", newMap)
	}

}

func (m *Module) OnDestroy() {
	log.Debug("center module destroy")
	for _, v := range Maps {
		v.closeSig <- true
	}
	m.wg.Wait()
}
