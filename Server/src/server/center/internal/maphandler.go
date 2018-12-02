package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"proto"
	"shared"
)

//处理地图初始化加载
func handleMapLoad(args []interface{}) {
	m := args[0].(*Map)
	log.Debug("handle map %v loading", m.Id)

	ChanRPC.Go("MapLoaded", m.Id)
}

//处理角色进入地图请求
func handleRoleEnterMap(args []interface{}) {
	m := args[0].(*Map)
	roleData := args[1].(*shared.RoleData)
	agent := args[2].(gate.Agent)

	rsp := &proto.RspEnterGs{}
	roleIdx := m.roleIndex.Alloc()
	if roleIdx < 0 {
		rsp.RetCode = 4 //地图人满
		agent.WriteMsg(rsp)
		return
	}

	//把角色绑定到Agent
	r := &MapRole{
		m:     m,
		data:  roleData,
		idx:   int32(roleIdx),
		agent: agent,
	}

	m.roles[roleIdx] = r
	agent.SetUserData(r)

	m.RoleEnter(agent)
}
