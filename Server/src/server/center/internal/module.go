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
	"time"
)

//等待角色进入项
type WaitEnterRole struct {
	roleId    string
	timeStamp int64
}

var (
	skeleton                            = base.NewSkeleton()
	ChanRPC                             = skeleton.ChanRPCServer
	waitEnterRoles                      = make(map[string]*WaitEnterRole)
	Maps                                = make(map[uint32]*Map)
	loadedMapCount                      = 0
	DBSession      *mongodb.DialContext = nil //数据库连接
	ctAgent        *TCPAgent
)

type Module struct {
	*module.Skeleton
	wgMap        sync.WaitGroup
	secondTicker *shared.UserTicker
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

		m.wgMap.Add(1)
		Maps[mapId] = newMap

		go func() {
			log.Debug("map %v goroutine start", mapId)
			newMap.Run(newMap.closeSig)
			log.Debug("map %v goroutine exit", mapId)
			m.wgMap.Done()
		}()

		newMap.ChanRPCServer.Go("LoadMap", newMap)
	}

	//开启计时器
	m.secondTicker = shared.NewUserTicker(time.Second, func() {
		for _, m := range Maps {
			m.ChanRPCServer.Go("OnTimerSecond", m)
		}
	})
}

func (m *Module) OnDestroy() {
	log.Debug("center module destroy")
	//停止计时器
	m.secondTicker.Stop()

	//停止地图协程
	for _, m := range Maps {
		m.closeSig <- true
	}

	m.wgMap.Wait()
}
