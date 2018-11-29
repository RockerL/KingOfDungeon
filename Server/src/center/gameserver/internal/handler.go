package internal

import (
	"center/conf"
	"center/gateclient"
	"github.com/name5566/leaf/gate"
	"proto"
	"reflect"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&proto.NotifyServerInited{}, handleServerInited)
}

//处理Game server 初始化完毕消息
func handleServerInited(args []interface{}) {
	r := args[0].(*proto.NotifyServerInited)
	_ = r

	a := args[1].(gate.Agent)

	runServers[a].isReg = true

	registeredServer++

	if registeredServer == conf.Server.ServerNum {
		gateclient.ChanRPC.Go("AllGameServerRegistered")
	}
}
