package internal

import (
	"center/base"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	"shared"
)

type GameServer struct {
	agent       gate.Agent
	playerCount int
	isReg       bool
	mapId       []uint32
}

var (
	skeleton         = base.NewSkeleton()
	ChanRPC          = skeleton.ChanRPCServer
	runServers       = make(map[gate.Agent]*GameServer) //已经连上来的游戏服务器
	registeredServer = 0
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	shared.GameServerChanRPC = ChanRPC
	log.Debug("game server module init")
}

func (m *Module) OnDestroy() {

}
