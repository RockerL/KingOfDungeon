package internal

import (
	"proto"
	"reflect"
)

func init() {
	handler(&proto.NotifyRoleEnter{}, handleRoleEnter)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

//处理中心服务器通知角色进入
func handleRoleEnter(args []interface{}) {

}
