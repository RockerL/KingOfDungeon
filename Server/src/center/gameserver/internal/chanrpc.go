package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

//一个game server 连上来
func rpcNewAgent(args []interface{}) {
	log.Debug("game server connected")

	a := args[0].(gate.Agent)

	runServers[a] = &GameServer{
		agent:       a,
		playerCount: 0,
		isReg:       false,
	}
}

//一个game server 断开
func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)

	if runServers[a].isReg {
		log.Debug("game server disconnected")
	}

	delete(runServers, a)
}
