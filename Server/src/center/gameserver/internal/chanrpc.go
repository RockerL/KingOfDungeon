package internal

import (
	"center/gameclient"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"github.com/rs/xid"
	"proto"
	"shared"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC("NotifyRoleEnter", rpcNotifyRoleEnter)
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

//用户选择角色进入game server
func rpcNotifyRoleEnter(args []interface{}) {
	user := args[0].(*shared.UserData)
	role := args[1].(*shared.RoleData)

	var agent gate.Agent = nil
	for k, v := range runServers {
		if !v.isReg {
			continue
		}

		for i := 0; i < len(v.mapId); i++ {
			if v.mapId[i] == role.MapId {
				agent = k
				break
			}
		}
	}

	roleId := role.Id.String()

	if agent != nil {
		agent.WriteMsg(&proto.NotifyRoleEnter{
			UserId: user.Id,
			RoleId: roleId,
			MapId:  role.MapId,
			Token:  xid.New().String(),
		})
	} else {
		gameclient.ChanRPC.Go("NotifyRoleEnteredReady", &proto.NotifyRoleEnteredReady{
			RetCode: 1,
			UserId:  user.Id,
			RoleId:  roleId,
		})
	}
}
