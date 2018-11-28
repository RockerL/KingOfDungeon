package shared

import (
	"github.com/name5566/leaf/network/protobuf"
	"proto"
)

var GSCTProccessor = protobuf.NewProcessor()

func init(){
	//game server to center server
	GSCTProccessor.Register(&proto.NotifyServerInited{})
	GSCTProccessor.Register(&proto.NotifyRoleEntered{})

	//center server to game server
	GSCTProccessor.Register(&proto.NotifyRoleEnter{})
}
