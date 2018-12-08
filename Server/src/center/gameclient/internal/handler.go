package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
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
	//检查运行商是否已经允许该用户登录

	//根据用户Id去查数据库，如果数据库没有，则新建一个用户到数据库
	s := dbSession.Ref()
	rsp := &proto.RspLogin{}
	defer func() {
		dbSession.UnRef(s)
		a.WriteMsg(rsp)
	}()

	//检查用户是否已经登录
	_, isUserLogin := loginUsers[r.UserId]
	if isUserLogin {
		rsp.RetCode = 1
		log.Debug("user %v is already login", r.UserId)
		return
	}

	user := &User{
		agent: a,
		state: Login,
	}
	err := s.DB(shared.DBName).C(shared.UserTableName).Find(bson.M{"Id": r.UserId}).One(&user.data)
	if err != nil {
		log.Debug("user %v err %v, create it", r.UserId, err.Error())
		user.data.RoleNum = 0
		user.data.Id = r.UserId
		s.DB(shared.DBName).C(shared.UserTableName).Insert(&user.data)
	}

	loginUsers[r.UserId] = user
	loginAgentUsers[a] = user

	rsp.RetCode = 0
	log.Debug("user id %v login successful as %v", r.UserId, a.RemoteAddr().String())
}

//处理获取角色列表
func handleRoleList(args []interface{}) {
	a := args[1].(gate.Agent)

	rsp := &proto.RspRolelist{}
	s := dbSession.Ref()
	defer func() {
		dbSession.UnRef(s)
		a.WriteMsg(rsp)
	}()

	user, ok := loginAgentUsers[a]
	if !ok {
		log.Debug("agent %v is not login", a.RemoteAddr())
		rsp.RetCode = 1 //用户未登录
		return
	}

	if user.agent != a {
		log.Debug("agent %v is not agent login %v", a.RemoteAddr(), user.agent.RemoteAddr())
		rsp.RetCode = 2 //错误的用户
		return
	}

	if user.state != Login {
		log.Debug("agent %v state wrong %v", a.RemoteAddr(), user.state)
		rsp.RetCode = 3 //错误的状态
		return
	}

	//去角色表里查找该用户的所有角色
	var roles []shared.RoleData
	s.DB(shared.DBName).C(shared.RoleTableName).Find(bson.M{"UserId": user.data.Id}).All(&roles)

	roleNum := len(roles)

	rsp.RetCode = 0
	rsp.RoleNum = user.data.RoleNum
	rsp.RoleInfos = make([]*proto.LoginRoleInfo, roleNum)
	for i := 0; i < roleNum; i++ {
		rsp.RoleInfos[i] = &proto.LoginRoleInfo{
			RoleId: roles[i].Id.String(),
			Name:   roles[i].Name,
			Sex:    roles[i].Sex,
			Level:  roles[i].Level,
			MapId:  roles[i].MapId,
		}
	}
}

//处理创建角色
func handleCreateRole(args []interface{}) {
	r := args[0].(*proto.ReqCreateRole)
	a := args[1].(gate.Agent)

	s := dbSession.Ref()
	rsp := &proto.RspCreateRole{}
	defer func() {
		dbSession.UnRef(s)
		a.WriteMsg(rsp)
	}()

	user, ok := loginAgentUsers[a]
	if !ok {
		rsp.RetCode = 1 //用户未登录
		return
	}

	//检查是否有空间创建角色
	if user.data.RoleNum >= shared.MaxRoleNum {
		rsp.RetCode = 2 //没有空间创建角色
		return
	}

	//检查是否角色名重名，长度是否过短
	if len(r.Name) < 4 {
		rsp.RetCode = 3 //角色名过短
		return
	}

	//创建角色数据
	role := shared.RoleData{
		Id:     bson.NewObjectId(),
		UserId: user.data.Id,
		Name:   r.Name,
		Sex:    r.Sex,
		Level:  1,
		MapId:  0, //=0表示随机进入某个地图某个位置
	}

	//写入数据库
	err := s.DB(shared.DBName).C(shared.RoleTableName).Insert(&role)
	if err != nil {
		log.Error("can not create role %v", err.Error())
		rsp.RetCode = 4 //数据库错误
		return
	}

	//返回消息给客户端
	rsp.Info.RoleId = role.Id.String()
	rsp.Info.Level = role.Level
	rsp.Info.Name = role.Name
	rsp.Info.Sex = role.Sex
}

//处理删除角色
func handleDelRole(args []interface{}) {
	r := args[0].(*proto.ReqDelRole)
	a := args[1].(gate.Agent)

	s := dbSession.Ref()
	rsp := &proto.RspDelRole{}
	defer func() {
		dbSession.UnRef(s)
		a.WriteMsg(rsp)
	}()

	_, ok := loginAgentUsers[a]
	if !ok {
		rsp.RetCode = 1 //用户未登录
		return
	}

	_id := bson.ObjectIdHex(r.RoleId)
	var role = new(shared.RoleData)
	err := s.DB(shared.DBName).C(shared.RoleTableName).Find(bson.M{"_id": _id}).One(&role)

	if err != nil {
		rsp.RetCode = 2 //没有找到角色数据
		return
	}

	s.DB(shared.DBName).C(shared.RoleTableName).RemoveAll(bson.M{"_id": _id})
	rsp.RetCode = 0
	rsp.RoleId = r.RoleId
}

//处理选择角色进入地图
func handleSelectRole(args []interface{}) {
	r := args[0].(*proto.ReqSelectRole)
	a := args[1].(gate.Agent)

	rsp := &proto.RspSelectRole{}
	s := dbSession.Ref()
	defer func() {
		dbSession.UnRef(s)
		if rsp.RetCode > 0 {
			a.WriteMsg(rsp)
		}
	}()

	user, ok := loginAgentUsers[a]
	if !ok {
		rsp.RetCode = 1 //用户未登录
		return
	}

	_id := bson.ObjectIdHex(r.RoleId)
	var role = new(shared.RoleData)
	err := s.DB(shared.DBName).C(shared.RoleTableName).Find(bson.M{"_id": _id}).One(&role)
	if err != nil {
		rsp.RetCode = 2 //无法查询到角色数据
		return
	}

	if user.data.Id != role.UserId {
		rsp.RetCode = 3 //角色数据和用户不对应
		return
	}

	//检查地图是否需要随机
	if role.MapId == 0 {
		role.MapId = 1
		role.Pos.X = rand.Float32() * shared.WorldSize
		role.Pos.Z = rand.Float32() * shared.WorldSize
		role.Pos.Y = 0

		err := s.DB(shared.DBName).C(shared.RoleTableName).Update(bson.M{"_id": _id}, role)
		if err != nil {
			rsp.RetCode = 4 //无法更新角色地图和位置
			return
		}
	}

	user.state = EnterGS

	//通知角色所在地图服务器
	shared.GameServerChanRPC.Go("NotifyRoleEnter", &user.data, role)
}
