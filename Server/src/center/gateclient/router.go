package gateclient

import (
	"center/gameclient"
	"center/msg"
	"proto"
)

func init() {
	//注册路由到game client模块的消息
	msg.CLProcessor.SetRouter(&proto.ReqLogin{}, gameclient.ChanRPC)
	msg.CLProcessor.SetRouter(&proto.ReqCreateRole{}, gameclient.ChanRPC)
	msg.CLProcessor.SetRouter(&proto.ReqDelRole{}, gameclient.ChanRPC)
	msg.CLProcessor.SetRouter(&proto.ReqRolelist{}, gameclient.ChanRPC)
	msg.CLProcessor.SetRouter(&proto.ReqSelectRole{}, gameclient.ChanRPC)
}
