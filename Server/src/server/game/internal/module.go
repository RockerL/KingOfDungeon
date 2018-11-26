package internal

import (
	"Common"
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/module"
	"server/base"
	"server/conf"
)

//运行时用户
type User struct {
	data Common.UserData
	selectRole string
	agent gate.Agent
}

//运行时角色
type Role struct {
	data Common.RoleData
	agent gate.Agent
}

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
	loginUsers = make(map[gate.Agent]*User)
	roleInfos = make(map[string]*Role)
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton

	//初始化地图
	for i := range conf.Server.MapLoad {
		fmt.Println(i)
	}
}

func (m *Module) OnDestroy() {

}
