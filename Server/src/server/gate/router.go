package gate

import (
	"proto"
	"server/center"
	"server/msg"
)

func init() {
	msg.CLProcessor.SetRouter(&proto.ReqEnterGs{}, center.ChanRPC)
	msg.CLProcessor.SetRouter(&proto.ReqRoleAction{}, center.ChanRPC)
}
