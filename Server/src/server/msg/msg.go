package msg

import (
	"github.com/name5566/leaf/network/protobuf"
	"server/conf"
)

//var Processor = json.NewProcessor()

var Processor = protobuf.NewProcessor()

func init() {
	Processor.SetByteOrder(conf.LittleEndian)
	Processor.Register(&RoleInfo{})
}

