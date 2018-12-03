package internal

import (
	"github.com/name5566/leaf/gate"
	"proto"
	"shared"
)

//运行时角色
type MapRole struct {
	m      *Map             //所在的地图
	c      *MapChunk        //地图区块
	data   *shared.RoleData //角色数据
	agent  gate.Agent       //角色对应网络客户端连接
	idx    int32            //角色在地图里的索引
	action uint32           //角色当前动作
}

func (r *MapRole) MakeBaseInfo() *proto.RoleBaseInfo {
	return &proto.RoleBaseInfo{
		Name:      r.data.Name,
		Level:     r.data.Level,
		LifePoint: r.data.LifePoint,
		Pos:       &proto.RolePos{X: r.data.Pos.X, Y: r.data.Pos.Y, Z: r.data.Pos.Z},
		RoleAngle: r.data.Angle,
		Outlook:   r.MakeOutlook(),
	}
}

func (r *MapRole) MakeOutlook() *proto.RoleOutlook {
	return &proto.RoleOutlook{
		Weapon: uint32(r.data.Outlook.Weapon),
		Helm:   uint32(r.data.Outlook.Helm),
		Face:   uint32(r.data.Outlook.Face),
		Wing:   uint32(r.data.Outlook.Wing),
		Bag:    uint32(r.data.Outlook.Bag),
		Suit:   uint32(r.data.Outlook.Suit),
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

	r.c = dst
	dst.AddRole(r)
}

func (r *MapRole) OnRoleLeave(role *MapRole) {
	leaveInfo := proto.SyncRoleLeaveRange{}
	leaveInfo.RoleIdx = r.idx
	role.agent.WriteMsg(&leaveInfo)

	leaveInfo.RoleIdx = role.idx
	r.agent.WriteMsg(&leaveInfo)
}

func (r *MapRole) OnRoleEnter(role *MapRole) {
	enterInfo := proto.SyncRoleEnterRange{}
	enterInfo.RoleIdx = r.idx
	enterInfo.RoleBaseInfo = r.MakeBaseInfo()
	role.agent.WriteMsg(&enterInfo)

	enterInfo.RoleIdx = role.idx
	enterInfo.RoleBaseInfo = role.MakeBaseInfo()
	r.agent.WriteMsg(&enterInfo)
}

func (r *MapRole) handleRoleAction(req *proto.ReqRoleAction) {
	//检查新的位置是否正确
	newChunkX := int32(req.Pos.X / shared.ChunkSize)
	newChunkZ := int32(req.Pos.Y / shared.ChunkSize)
	newChunk := r.m.GetChunk(newChunkX, newChunkZ)
	if newChunk == nil {
		return
	}

	//检查是否要切换区域
	if newChunk != r.c {
		r.ChangeChunk(r.c, newChunk)
	}

	r.data.Pos.X = req.Pos.X
	r.data.Pos.Y = req.Pos.Y
	r.data.Pos.Z = req.Pos.Z
	r.data.Angle = req.RoleAngle
	r.action = req.Action

	rsp := &proto.SyncRoleAction{
		Pos:       req.Pos,
		RoleAngle: req.RoleAngle,
		Action:    req.Action,
	}

	r.m.BroadcastAroundRole(r, rsp)
}
