package internal

import (
	"github.com/name5566/leaf/gate"
	"gopkg.in/mgo.v2/bson"
	"proto"
	"reflect"
	"shared"
	"time"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&proto.ReqEnterGs{}, handleEnterGS)
}

//处理客户端请求进入游戏服务器
func handleEnterGS(args []interface{}) {
	r := args[0].(*proto.ReqEnterGs)
	a := args[1].(gate.Agent)

	rsp := &proto.RspEnterGs{}

	s := DBSession.Ref()
	defer DBSession.UnRef(s)

	curTime := time.Now().Unix()

	//该agent已经关联了地图角色
	if a.UserData() != nil {
		rsp.RetCode = 1
		a.WriteMsg(rsp)
		return
	}

	//检查过期的令牌，令牌存在的时间为30秒
	for k, v := range waitEnterRoles {
		if curTime-v > 30 {
			delete(waitEnterRoles, k)
		}
	}

	//检查令牌有效性
	_, ok := waitEnterRoles[r.Token]
	if !ok {
		rsp.RetCode = 2
		a.WriteMsg(rsp)
		return
	}

	_id := bson.ObjectIdHex(r.RoleId)
	var role = new(shared.RoleData)
	err := s.DB(shared.DBName).C(shared.RoleTableName).Find(bson.M{"_id": _id}).One(&role)
	if err != nil {
		rsp.RetCode = 3
		a.WriteMsg(rsp)
		return
	}

	//发消息给地图协程
	m, ok := Maps[role.MapId]
	m.ChanRPCServer.Go("RoleEnterMap", m, role, a)
}
