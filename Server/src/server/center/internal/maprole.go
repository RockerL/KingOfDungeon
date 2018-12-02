package internal

import (
	"github.com/name5566/leaf/gate"
	"proto"
	"shared"
)

//运行时角色
type MapRole struct {
	m     *Map             //所在的地图
	data  *shared.RoleData //角色数据
	agent gate.Agent       //角色对应网络客户端连接
	idx   int32            //角色在地图里的索引
}

func (r *MapRole) MakeBaseInfo() *proto.RoleBaseInfo {
	return &proto.RoleBaseInfo{
		Name:  r.data.Name,
		Level: r.data.Level,
	}
}
