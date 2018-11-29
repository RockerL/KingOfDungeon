package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2/bson"
	"proto"
	"reflect"
	"shared"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&proto.ReqLogin{}, handleLogin)
	handleMsg(&proto.ReqCreateRole{}, handleCreateRole)
	handleMsg(&proto.ReqRolelist{}, handleRoleList)
	handleMsg(&proto.ReqDelRole{}, handleDelRole)
	handleMsg(&proto.ReqSelectRole{}, handleSelectRole)
}

//处理登录请求
func handleLogin(args []interface{}) {
	r := args[0].(*proto.ReqLogin)
	a := args[1].(gate.Agent)

	_, isUserLogin := loginUsers[r.UserId]
	if isUserLogin {
		a.WriteMsg(&proto.RspLogin{
			RetCode: 1, //用户已经登录
		})
		log.Debug("user %v is already login", r.UserId)
		return
	}

	//检查运行商是否已经允许该用户登录

	//根据用户Id去查数据库，如果数据库没有，则新建一个用户到数据库
	s := dbSession.Ref()
	defer dbSession.UnRef(s)
	var users []shared.UserData
	s.DB(DBName).C(UserTableName).Find(bson.M{"OpId": r.UserId}).One(&users)

	if len(users) == 0 {
		log.Debug("user %v is not found, create it", r.UserId)

		s.DB(DBName).C(UserTableName).Insert(&shared.UserData{
			Id:      r.UserId,
			RoleNum: 0,
		})
	}

	s.DB(DBName).C(UserTableName).Find(bson.M{"OpId": r.UserId}).One(&users)

	user := &User{
		data:  users[0],
		agent: a,
	}

	loginUsers[r.UserId] = user

	loginAgentUsers[a] = user

	a.WriteMsg(&proto.RspLogin{
		RetCode: 0,
	})

	log.Debug("user id %v login successful as %v", r.UserId, a.RemoteAddr())
}

//处理获取角色列表
func handleRoleList(args []interface{}) {
	a := args[1].(gate.Agent)

	user, ok := loginAgentUsers[a]
	if !ok {
		log.Debug("agent %v is not login", a.RemoteAddr())
		a.WriteMsg(&proto.RspRolelist{
			RetCode: 1, //用户未登录
			RoleNum: 0,
		})
		return
	}

	if user.agent != a {
		log.Debug("agent %v is not agent login %v", a.RemoteAddr(), user.agent.RemoteAddr())
		a.WriteMsg(&proto.RspRolelist{
			RetCode: 2, //错误的用户
			RoleNum: 0,
		})
		return
	}

	//去角色表里查找该用户的所有角色
	s := dbSession.Ref()
	defer dbSession.UnRef(s)
	var roles []shared.RoleData
	s.DB(DBName).C(RoleTableName).Find(bson.M{"UserId": user.data.Id}).All(&roles)

	roleNum := len(roles)

	retMsg := proto.RspRolelist{
		RetCode:   0,
		RoleNum:   user.data.RoleNum,
		RoleInfos: make([]*proto.LoginRoleInfo, roleNum),
	}

	for i := 0; i < roleNum; i++ {
		retMsg.RoleInfos[i] = &proto.LoginRoleInfo{
			RoleId: roles[i].Id.String(),
			Name:   roles[i].Name,
			Sex:    roles[i].Sex,
			Level:  roles[i].Level,
			MapId:  roles[i].MapId,
		}
	}

	a.WriteMsg(&retMsg)
}

//处理创建角色
func handleCreateRole(args []interface{}) {
	r := args[0].(*proto.ReqCreateRole)
	a := args[1].(gate.Agent)

	user, ok := loginAgentUsers[a]
	if !ok {
		a.WriteMsg(&proto.RspCreateRole{
			RetCode: 1, //用户未登录
		})
		return
	}

	//检查是否有空间创建角色
	if user.data.RoleNum >= MaxRoleNum {
		a.WriteMsg(&proto.RspCreateRole{
			RetCode: 2, //没有空间创建角色
		})
		return
	}

	//检查是否角色名重名，长度是否过短
	if len(r.Name) < 4 {
		a.WriteMsg(&proto.RspCreateRole{
			RetCode: 3, //角色名过短
		})
		return
	}

	//创建角色数据
	role := shared.RoleData{
		Id:     bson.NewObjectId(),
		UserId: user.data.Id,
		Name:   r.Name,
		Sex:    r.Sex,
		Level:  1,
		MapId:  0,
	}

	//写入数据库
	s := dbSession.Ref()
	defer dbSession.UnRef(s)
	err := s.DB(DBName).C(RoleTableName).Insert(&role)
	if err != nil {
		log.Error("can not create role %v", err.Error())
		a.WriteMsg(&proto.RspCreateRole{
			RetCode: 3, //数据库错误
		})
		return
	}

	//返回消息给客户端
	a.WriteMsg(&proto.RspCreateRole{
		RetCode: 0,
		Info: &proto.LoginRoleInfo{
			RoleId: role.Id.String(),
			Name:   role.Name,
			Sex:    role.Sex,
			Level:  role.Level,
		},
	})
}

//处理删除角色
func handleDelRole(args []interface{}) {
	r := args[0].(*proto.ReqDelRole)
	a := args[1].(gate.Agent)

	_, ok := loginAgentUsers[a]
	if !ok {
		a.WriteMsg(&proto.RspDelRole{
			RetCode: 1, //用户未登录
			RoleId:  r.RoleId,
		})
		return
	}

	s := dbSession.Ref()
	defer dbSession.UnRef(s)

	_id := bson.ObjectIdHex(r.RoleId)
	var role = new(shared.RoleData)
	err := s.DB(DBName).C(RoleTableName).Find(bson.M{"_id": _id}).One(&role)

	if err != nil {
		a.WriteMsg(&proto.RspDelRole{
			RetCode: 2, //没有找到角色数据
			RoleId:  r.RoleId,
		})
		return
	}

	s.DB(DBName).C(RoleTableName).RemoveAll(bson.M{"_id": _id})

	a.WriteMsg(&proto.RspDelRole{
		RetCode: 0,
		RoleId:  r.RoleId,
	})
}

//处理选择角色进入地图
func handleSelectRole(args []interface{}) {
	r := args[0].(*proto.ReqSelectRole)
	a := args[1].(gate.Agent)

	user, ok := loginAgentUsers[a]
	if !ok {
		a.WriteMsg(&proto.RspSelectRole{
			RetCode: 1, //用户未登录
			RoleId:  r.RoleId,
		})
		return
	}

	s := dbSession.Ref()
	defer dbSession.UnRef(s)

	_id := bson.ObjectIdHex(r.RoleId)
	var role = new(shared.RoleData)
	err := s.DB(DBName).C(RoleTableName).Find(bson.M{"_id": _id}).One(&role)
	if err != nil {
		a.WriteMsg(&proto.RspSelectRole{
			RetCode: 2, //没有找到角色数据
			RoleId:  r.RoleId,
		})
		return
	}

	if user.data.Id != role.Id.String() {
		a.WriteMsg(&proto.RspSelectRole{
			RetCode: 3, //角色数据和用户不对应
			RoleId:  r.RoleId,
		})
		return
	}

	//通知角色所在地图服务器
	shared.GameServerChanRPC.Go("NotifyRoleEnter", user)
}
