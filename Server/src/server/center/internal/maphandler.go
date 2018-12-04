package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"proto"
	"shared"
)

//运行在地图协程里的函数

//处理地图初始化加载
func mapHandleMapLoad(args []interface{}) {
	m := args[0].(*Map)
	log.Debug("handle map %v loading", m.Id)

	ChanRPC.Go("MapLoaded", m.Id)
}

//处理角色进入地图请求
func mapHandleRoleEnterMap(args []interface{}) {
	a := args[0].(gate.Agent)
	roleData := args[1].(*shared.RoleData)
	m := args[2].(*Map)

	//如果角色在离开等待列表里则直接使用已有的角色，并替换成新的Agent
	var role *MapRole = nil
	for k, v := range m.waitLeaveRoles {
		if v.role.data.Id == roleData.Id {
			role = v.role
			role.agent = a
			delete(m.waitLeaveRoles, k)
			break
		}
	}

	if role == nil {
		rsp := &proto.RspEnterGs{}
		roleIdx := m.roleIdxAllocator.Alloc()
		if roleIdx < 0 {
			rsp.RetCode = 5 //地图人满
			a.WriteMsg(rsp)
			return
		}

		role = &MapRole{
			m:     m,
			data:  roleData,
			idx:   int32(roleIdx),
			agent: a,
		}
	}

	a.SetUserData(role) //把角色绑定到Agent
	m.RoleEnter(role)
}

//客户端断线，5秒后移除出地图
func mapHandleRoleDisconnect(args []interface{}) {
	role := args[0].(*MapRole)
	role.agent.SetUserData(nil)
	role.agent = nil

	role.m.waitLeaveRoles[role.data.Id.String()] = &WaitLeaveRole{role, 5}
}

//处理角色位置，方向，动作
func mapHandleRoleAction(args []interface{}) {
	r := args[0].(*proto.ReqRoleAction)
	a := args[1].(gate.Agent)
	role := a.UserData().(*MapRole)
	role.handleRoleAction(r)
}

//每1秒调用
func mapHandleOnTimerSecond(args []interface{}) {
	m := args[0].(*Map)
	//log.Debug("On timer %v", m.Id)
	for k, v := range m.waitLeaveRoles {
		v.remainSecond--
		if v.remainSecond <= 0 {
			delete(m.waitLeaveRoles, k)
			m.RoleLeave(v.role)
		}
	}
}
