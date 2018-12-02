package internal

import (
	"github.com/name5566/leaf/gate"
	"proto"
	"shared"
)

//运行时角色
type MapRole struct {
	Next  *MapRole         //下一个节点
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

func (r *MapRole) ChangeChunk(src *MapChunk, dst *MapChunk) {
	if src == dst {
		return
	}
	//搜集离开区域的周围区域（含）以及进入区域的周围区域
	var leaveChunks [shared.ClientChunkTotal]*MapChunk
	var enterChunks [shared.ClientChunkTotal]*MapChunk
	if src != nil {
		idx := 0
		startChunkX := src.data.ChunkX - int32(shared.ClientChunkNum)/2
		startChunkZ := src.data.ChunkZ - int32(shared.ClientChunkNum)/2
		for z := startChunkZ; z < startChunkZ+shared.ClientChunkNum; z++ {
			for x := startChunkX; x < startChunkX+shared.ClientChunkNum; x++ {
				leaveChunks[idx] = r.m.GetChunk(x, z)
				idx++
			}
		}
	}

	if dst != nil {
		idx := 0
		startChunkX := dst.data.ChunkX - int32(shared.ClientChunkNum)/2
		startChunkZ := dst.data.ChunkZ - int32(shared.ClientChunkNum)/2
		for z := startChunkZ; z < startChunkZ+shared.ClientChunkNum; z++ {
			for x := startChunkX; x < startChunkX+shared.ClientChunkNum; x++ {
				enterChunks[idx] = r.m.GetChunk(x, z)
				idx++
			}
		}
	}

	src.RemoveRole(r)

	for i := 0; i < shared.ClientChunkTotal; i++ {
		c := leaveChunks[i]
		if c == nil {
			continue
		}
		isFound := false
		for j := 0; j < shared.ClientChunkTotal; j++ {
			if c == enterChunks[j] {
				isFound = true
				break
			}
		}
		if isFound {
			continue
		}
		c.OnRoleLeave(r)
	}

	for i := 0; i < shared.ClientChunkTotal; i++ {
		c := enterChunks[i]
		if c == nil {
			continue
		}
		isFound := false
		for j := 0; j < shared.ClientChunkTotal; j++ {
			if c == leaveChunks[j] {
				isFound = true
				break
			}
		}
		if isFound {
			continue
		}
		c.OnRoleEnter(r)
	}

	dst.AddRole(r)
}
