package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"proto"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC("NotifyRoleEnteredReady", rpcNotifyRoleEnteredReady)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)

	log.Debug("agent create %v", a.RemoteAddr())

}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)

	user, ok := loginAgentUsers[a]
	if ok {
		log.Debug("agent close %v", a.RemoteAddr())
		delete(loginUsers, user.data.Id)
		delete(loginAgentUsers, a)
	}
}

//收到通知game server准备好让玩家进入
func rpcNotifyRoleEnteredReady(args []interface{}) {
	rsp := args[0].(*proto.NotifyRoleEnteredReady)

	user, ok := loginUsers[rsp.UserId]
	if ok {
		user.agent.WriteMsg(&proto.RspSelectRole{
			RetCode:    rsp.RetCode,
			RoleId:     rsp.RoleId,
			Token:      rsp.Token,
			MapId:      rsp.MapId,
			ServerIp:   rsp.ServerIp,
			ServerPort: rsp.ServerPort,
		})
	}
}
