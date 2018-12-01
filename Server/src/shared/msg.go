package shared

import (
	"github.com/name5566/leaf/network/protobuf"
	"proto"
)

var GSCTProcessor = protobuf.NewProcessor()

func init() {
	//game server to center server
	GSCTProcessor.Register(&proto.NotifyServerInited{})
	GSCTProcessor.Register(&proto.NotifyRoleEnteredReady{})

	//center server to game server
	GSCTProcessor.Register(&proto.NotifyRoleEnter{})
}
