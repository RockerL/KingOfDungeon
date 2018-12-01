package msg

import (
	"github.com/name5566/leaf/network/protobuf"
	"proto"
	"server/conf"
)

//var Processor = json.NewProcessor()

var CLProcessor = protobuf.NewProcessor()

func init() {
	CLProcessor.SetByteOrder(conf.LittleEndian)

	//client to game server
	CLProcessor.Register(&proto.ReqEnterGs{})

	//game server to client
	CLProcessor.Register(&proto.RspEnterGs{})
}
