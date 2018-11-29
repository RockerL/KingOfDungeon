package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)

	log.Debug("agent create %v", a.RemoteAddr())

}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)

	_, ok := loginAgentUsers[a]
	if ok {
		log.Debug("agent close %v", a.RemoteAddr())
		delete(loginAgentUsers, a)
	}

}
