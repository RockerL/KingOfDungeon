package msg

import (
	"center/conf"
	"github.com/name5566/leaf/network/protobuf"

)

//var Processor = json.NewProcessor()

var Processor = protobuf.NewProcessor()

func init() {
	Processor.SetByteOrder(conf.LittleEndian)

	//用反射注册消息类型
	Processor.Register(&ReqLogin{})
	Processor.Register(&ReqCreateRole{})
	Processor.Register(&ReqSelectRole{})
	Processor.Register(&ReqRolelist{})
	Processor.Register(&ReqDelRole{})
}

