package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"proto"
)

//处理地图初始化加载
func mapHandleMapLoad(args []interface{}) {
	m := args[0].(*Map)
	log.Debug("handle map %v loading", m.Id)

	ChanRPC.Go("MapLoaded", m.Id)
}

//处理角色进入地图请求
func mapHandleRoleEnterMap(args []interface{}) {
	a := args[0].(gate.Agent)
	m := a.UserData().(*MapRole).m
	m.RoleEnter(a)
}

//处理角色位置，方向，动作
func mapHandleRoleAction(args []interface{}) {
	r := args[0].(*proto.ReqRoleAction)
	a := args[1].(gate.Agent)
	role := a.UserData().(*MapRole)
	role.handleRoleAction(r)
}
