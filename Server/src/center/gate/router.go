package gate

import (
	"center/login"
	"center/msg"
)

func init() {
	//注册路由到login模块的消息
	msg.Processor.SetRouter(&msg.ReqLogin{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.ReqCreateRole{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.ReqDelRole{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.ReqRolelist{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.ReqSelectRole{}, login.ChanRPC)

	//注册路由到game模块的消息
}
