package internal

import (
	"center/conf"
	"center/gateclient"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	num := len(runServers)
	if num >= conf.Server.ServerNum {
		return
	}

	log.Debug("new game server registered")

	a := args[0].(gate.Agent)

	runServers[a] = &GameServer{
		agent:a,
		playerCount:0,
	}

	if len(runServers) == conf.Server.ServerNum {
		gateclient.ChanRPC.Go("GameServerRegistered")
	}
}

func rpcCloseAgent(args []interface{}) {
	//a := args[0].(gate.Agent)
}