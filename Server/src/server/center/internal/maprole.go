package internal

import (
	"github.com/name5566/leaf/gate"
	"proto"
	"shared"
	"time"
)

//运行时角色
type MapRole struct {
	m      *Map             //所在的地图
	c      *MapChunk        //地图区块
	data   *shared.RoleData //角色数据
	agent  gate.Agent       //角色对应网络客户端连接
	idx    int32            //角色在地图里的索引，地图里唯一
	action uint32           //角色当前动作
	face   uint16           //当前表情
	opTime int64            //上一次操作的时间戳
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
		Weapon: uint32(r.data.EquipData.Weapon.Type),
		Helm:   uint32(r.data.EquipData.Helm.Type),
		Face:   uint32(r.face),
		Wing:   uint32(r.data.EquipData.Wing.Type),
		Bag:    uint32(r.data.EquipData.Bag.Type),
		Suit:   uint32(r.data.EquipData.Suit.Type),
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

	if src != nil {
		src.RemoveRole(r)
	}

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

		syncLeave := &proto.SyncChunkLeaveRange{
			ChunkX: c.data.ChunkX,
			ChunkZ: c.data.ChunkZ,
		}
		if r.agent != nil {
			r.agent.WriteMsg(syncLeave)
		}
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

		syncEnter := &proto.SyncChunkEnterRange{
			ChunkX: c.data.ChunkX,
			ChunkZ: c.data.ChunkZ,
			Blocks: c.MakeChunkBlockInfo(),
		}
		if r.agent != nil {
			r.agent.WriteMsg(syncEnter)
		}
	}

	r.c = dst
	if dst != nil {
		dst.AddRole(r)
	}
}

func (r *MapRole) OnRoleLeave(role *MapRole) {
	leaveInfo := proto.SyncRoleLeaveRange{}
	leaveInfo.RoleIdx = r.idx
	if role.agent != nil {
		role.agent.WriteMsg(&leaveInfo)
	}

	leaveInfo.RoleIdx = role.idx
	if role.agent != nil {
		r.agent.WriteMsg(&leaveInfo)
	}
}

func (r *MapRole) OnRoleEnter(role *MapRole) {
	enterInfo := proto.SyncRoleEnterRange{}
	enterInfo.RoleIdx = r.idx
	enterInfo.RoleBaseInfo = r.MakeBaseInfo()
	if role.agent != nil {
		role.agent.WriteMsg(&enterInfo)
	}

	enterInfo.RoleIdx = role.idx
	enterInfo.RoleBaseInfo = role.MakeBaseInfo()
	if r.agent != nil {
		r.agent.WriteMsg(&enterInfo)
	}
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

	r.m.BroadcastAroundRole(r.c, rsp)
}

func (r *MapRole) handleRoleOpBlock(req *proto.ReqOpBlock) {
	chunkX := int32(req.BlockX) / shared.ChunkBlockNum
	chunkZ := int32(req.BlockZ) / shared.ChunkBlockNum

	chunk := r.m.GetChunk(chunkX, chunkZ)
	if chunk == nil {
		return
	}

	block := chunk.GetBlock(req.BlockX, req.BlockY, req.BlockZ)
	if block == nil {
		return
	}

	roleBlockX := int32(r.data.Pos.X / shared.BlockSize)
	roleBlockZ := int32(r.data.Pos.Z / shared.BlockSize)

	//角色只能操作自己周围的块
	if shared.Int32Abs(roleBlockX-int32(req.BlockX)) > 1 ||
		shared.Int32Abs(roleBlockZ-int32(req.BlockZ)) > 1 {
		return
	}

	//检查CD
	curTime := time.Now().Unix()
	if curTime-r.opTime < 1 {
		return
	}

	sync := &proto.SyncBlock{
		BlockX: req.BlockX,
		BlockY: req.BlockY,
		BlockZ: req.BlockZ,
	}

	switch req.OpCode {
	case OpDig:
		if block.Content > 0 {
			block.Content--
		}
		if block.Content == 0 {
			block.BlockType = BlockAir
		}
	}

	sync.Info = &proto.BlockInfo{
		BlockType: uint32(block.BlockType),
		Content:   uint32(block.Content),
	}
	r.m.BroadcastAroundRole(chunk, sync)
}
