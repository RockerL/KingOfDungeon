package internal

import (
	"Common"
	"center/base"
	"center/conf"
	"github.com/name5566/leaf/db/mongodb"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
)

const DBName = "game"
const UserTableName = "user"
const RoleTableName = "role"
const MaxRoleNum = 8

//运行时用户
type User struct {
	data       Common.UserData
	selectRole string
	agent      gate.Agent
}

var (
	skeleton   = base.NewSkeleton()				//新建框架实例,skeleton 实现了 Module 接口的 Run 方法并提供了ChanRPC goroutine 定时器
	ChanRPC    = skeleton.ChanRPCServer			//导出给外界使用
	loginUsers = make(map[gate.Agent]*User)		//已经登录成功的在线用户
	dbSession  *mongodb.DialContext = nil		//数据库连接
)

type Module struct {
	*module.Skeleton		//组合了Skeleton模块，实现了Module
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton

	//初始化数据连接
	dbSession,err := mongodb.Dial(conf.Server.DBAddr, 10)
	if dbSession == nil {
		log.Release("can not connect mongodb ip %v err %v", conf.Server.DBAddr, err.Error())
		return
	} else {
		log.Release("connect mongodb %v success", conf.Server.DBAddr)

		log.Debug("game client module init")
	}
}

func (m *Module) OnDestroy() {
	if dbSession != nil {
		log.Release("mongodb closed")
		dbSession.Close()
	}
}
