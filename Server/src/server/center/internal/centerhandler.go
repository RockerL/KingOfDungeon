package internal

import (
	"proto"
	"reflect"
	"server/conf"
	"shared/netutil"
	"time"
)

func init() {
	handler(&proto.NotifyRoleEnter{}, handleNotifyRoleEnter)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

//处理中心服务器通知角色进入
func handleNotifyRoleEnter(args []interface{}) {
	a := args[0].(*proto.NotifyRoleEnter)

	waitEnterRoles[a.Token] = &WaitEnterRole{
		roleId:    a.RoleId,
		timeStamp: time.Now().Unix(),
	}

	ctAgent.WriteMsg(&proto.NotifyRoleEnteredReady{
		RetCode:    0,
		UserId:     a.UserId,
		RoleId:     a.RoleId,
		MapId:      a.MapId,
		Token:      a.Token,
		ServerIp:   netutil.GetIPFromAddr(conf.Server.TCPAddr),
		ServerPort: netutil.GetPortFromAddr(conf.Server.TCPAddr),
	})
}
