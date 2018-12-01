package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"proto"
	"shared"
)

func handleMapLoad(args []interface{}) {
	m := args[0].(*Map)
	log.Debug("handle map %v loading", m.Id)

	ChanRPC.Go("MapLoaded", m.Id)
}

func handleRoleEnterMap(args []interface{}) {
	m := args[0].(*Map)
	roleData := args[1].(*shared.RoleData)
	agent := args[2].(gate.Agent)

	rsp := &proto.RspEnterGs{}
	roleIdx := m.roleIndex.Alloc()
	if roleIdx < 0 {
		rsp.RetCode = 4
		agent.WriteMsg(rsp)
		return
	}

	r := &MapRole{
		m:     m,
		data:  roleData,
		idx:   roleIdx,
		agent: agent,
	}

	m.roles[roleIdx] = r

	agent.SetUserData(r)

	rsp.RetCode = 0
	agent.WriteMsg(rsp)
}
