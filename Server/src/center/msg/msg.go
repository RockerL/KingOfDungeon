package msg

import (
	"center/conf"
	"github.com/name5566/leaf/network/protobuf"
	"proto"
)

//var Processor = json.NewProcessor()

var CLProcessor = protobuf.NewProcessor()

func init() {
	CLProcessor.SetByteOrder(conf.LittleEndian)

	//client to center server
	CLProcessor.Register(&proto.ReqLogin{})
	CLProcessor.Register(&proto.ReqCreateRole{})
	CLProcessor.Register(&proto.ReqSelectRole{})
	CLProcessor.Register(&proto.ReqRolelist{})
	CLProcessor.Register(&proto.ReqDelRole{})

	//center server to client
	CLProcessor.Register(&proto.RspLogin{})
	CLProcessor.Register(&proto.RspCreateRole{})
	CLProcessor.Register(&proto.RspSelectRole{})
	CLProcessor.Register(&proto.RspRolelist{})
	CLProcessor.Register(&proto.RspDelRole{})

}

