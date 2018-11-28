package gate

import (
	"proto"
	"server/client"
	"server/msg"
)

func init() {
	msg.CLProcessor.SetRouter(&proto.ReqGetRole{}, client.ChanRPC)
}
