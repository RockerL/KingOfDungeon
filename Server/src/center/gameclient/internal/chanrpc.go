package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"server/game"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)

	log.Debug("agent create %v", a.RemoteAddr())

	game.ChanRPC.Go("NewAgent", a)
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)

	_, ok := loginUsers[a]
	if ok {
		log.Debug("agent close %v", a.RemoteAddr())
		delete(loginUsers, a)
	}

	game.ChanRPC.Go("CloseAgent", a)
}
