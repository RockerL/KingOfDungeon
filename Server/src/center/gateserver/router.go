package gateserver

import (
	"center/gameserver"
	"proto"
	"shared"
)

func init() {
	//注册路由到gameserver模块的消息
	shared.GSCTProcessor.SetRouter(&proto.NotifyServerInited{}, gameserver.ChanRPC)
}
