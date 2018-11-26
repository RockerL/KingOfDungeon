package gateclient

import (
	"center/gameclient"
	"center/msg"
)

func init() {
	//注册路由到login模块的消息
	msg.Processor.SetRouter(&msg.ReqLogin{}, gameclient.ChanRPC)
	msg.Processor.SetRouter(&msg.ReqCreateRole{}, gameclient.ChanRPC)
	msg.Processor.SetRouter(&msg.ReqDelRole{}, gameclient.ChanRPC)
	msg.Processor.SetRouter(&msg.ReqRolelist{}, gameclient.ChanRPC)
	msg.Processor.SetRouter(&msg.ReqSelectRole{}, gameclient.ChanRPC)

	//注册路由到game模块的消息
}
