package internal

import (
	"center/base"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
)

type GameServer struct {
	agent gate.Agent
	playerCount int
}

var (
	skeleton   = base.NewSkeleton()
	ChanRPC    = skeleton.ChanRPCServer
	runServers = make(map[gate.Agent]*GameServer)		//已经连上来的游戏服务器
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton

	log.Debug("game server module init")
}

func (m *Module) OnDestroy() {

}
