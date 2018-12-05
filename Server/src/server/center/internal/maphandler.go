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

	//如果角色在离开等待列表里则从等待列表里删除
	for k, v := range m.waitLeaveRoles {
		if v.role.data.Id == roleData.Id {
			delete(m.waitLeaveRoles, k)
			break
		}
	}

	//如果角色正在玩，关闭之前的旧链接，关联新链接
	role, ok := m.rolesMap[roleData.Id]
	if ok {
		if role.agent != nil {
			role.agent.SetUserData(nil)
			role.agent.Destroy()
		}
		role.agent = a //关联新链接
	} else {
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

//处理角色操作块
func mapHandleRoleOpBlock(args []interface{}) {
	r := args[0].(*proto.ReqOpBlock)
	a := args[1].(gate.Agent)
	role := a.UserData().(*MapRole)
	role.handleRoleOpBlock(r)
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
