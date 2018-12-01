package shared

import "gopkg.in/mgo.v2/bson"

//用户信息，可以序列化到json
type UserData struct {
	Id      string `bson:"Id"`      //运营商的用户ID
	RoleNum uint32 `bson:"RoleNum"` //角色数量
}

//角色信息，可以序列化到json
type RoleData struct {
	Id     bson.ObjectId `bson:"_id"`    //数据库ID
	UserId string        `bson:"UserId"` //角色对应的用户ID
	Name   string        `bson:"Name"`   //角色名字
	Sex    uint32        `bson:"Sex"`    //性别
	Level  uint32        `bson:"Level"`  //等级
	MapId  uint32        `bson:"MapId"`  //所在地图ID
}

const DBName = "game"        //保存用户数据和角色数据
const UserTableName = "user" //用户数据表名
const RoleTableName = "role" //角色数据表名
const MaxRoleNum = 8
