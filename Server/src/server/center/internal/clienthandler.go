package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
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
	handleMsg(&proto.ReqRoleAction{}, handleRoleAction)
}

//处理客户端请求进入游戏服务器
func handleEnterGS(args []interface{}) {
	r := args[0].(*proto.ReqEnterGs)
	a := args[1].(gate.Agent)

	rsp := &proto.RspEnterGs{}

	//该agent已经关联了地图角色
	if a.UserData() != nil {
		rsp.RetCode = 1
		a.WriteMsg(rsp)
		return
	}

	//检查过期的令牌，令牌存在的时间为30秒
	curTime := time.Now().Unix()
	for k, v := range waitEnterRoles {
		if curTime-v.timeStamp > 30 {
			delete(waitEnterRoles, k)
		}
	}

	//检查令牌有效性
	waitRole, ok := waitEnterRoles[r.Token]
	if !ok || waitRole.roleId != r.RoleId {
		rsp.RetCode = 2 //没有令牌，或者令牌对应的角色不符
		a.WriteMsg(rsp)
		return
	}

	delete(waitEnterRoles, r.Token)

	s := DBSession.Ref()
	defer DBSession.UnRef(s)
	_id := bson.ObjectIdHex(r.RoleId)
	var roleData = new(shared.RoleData)
	err := s.DB(shared.DBName).C(shared.RoleTableName).Find(bson.M{"_id": _id}).One(&roleData)
	if err != nil {
		rsp.RetCode = 3 //角色数据不存在
		a.WriteMsg(rsp)
		return
	}

	//发消息给角色所在的地图协程
	m, ok := Maps[roleData.MapId]
	if ok {
		m.ChanRPCServer.Go("RoleEnterMap", a, roleData, m)
	} else {
		log.Error("role want enter map %v, but not in this server", roleData.MapId)
		rsp.RetCode = 4
		a.WriteMsg(rsp)
		return
	}
}

//处理角色的动作，位置，方向同步
func handleRoleAction(args []interface{}) {
	a := args[1].(gate.Agent)
	mapRole := a.UserData().(*MapRole)
	if mapRole != nil {
		mapRole.m.ChanRPCServer.Go("RoleAction", args)
	}
}
