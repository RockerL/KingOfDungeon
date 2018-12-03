package internal

import (
	"proto"
	"shared"
	"shared/algorithm"
)

//运行时地图块
type MapChunk struct {
	data  MapChunkData          //地图区块数据
	roles *algorithm.SingleList //角色的链表头
}

//按照大平地生成地形
func (c *MapChunk) InitChunkWithFlat(id string, chunkX int32, chunkZ int32) {
	c.data.ChunkId = id
	c.data.ChunkX = chunkX
	c.data.ChunkZ = chunkZ
	for y := 0; y < shared.BlockMaxY; y++ {
		for z := 0; z < shared.ChunkBlockNum; z++ {
			for x := 0; x < shared.ChunkBlockNum; x++ {
				t := 0
				if y < shared.BlockMaxY/2 {
					t = Earth
				} else {
					t = Air
				}
				idx := y*shared.ChunkBlockNum*shared.ChunkBlockNum + z*shared.ChunkBlockNum + x
				c.data.BlockArray[idx].BlockType = uint8(t)
			}
		}
	}
}

func (c *MapChunk) MakeChunkInfo() *proto.ChunkInfo {
	info := &proto.ChunkInfo{
		Blocks: make([]*proto.BlockInfo, len(c.data.BlockArray)),
	}

	for i := 0; i < len(c.data.BlockArray); i++ {
		info.Blocks[i] = &proto.BlockInfo{
			BlockType: uint32(c.data.BlockArray[i].BlockType),
			SubType:   uint32(c.data.BlockArray[i].SubType),
			Durable:   c.data.BlockArray[i].Durable,
		}
	}
	return info
}

func (c *MapChunk) RemoveRole(role *MapRole) bool {
	return c.roles.Delete(role)
}

func (c *MapChunk) AddRole(role *MapRole) {
	c.roles.Insert(role)
}

func (c *MapChunk) OnRoleLeave(role *MapRole) {
	//遍历角色同步离开信息
	c.roles.Traversal(func(v algorithm.ElemType) {
		r := v.(*MapRole)
		role.OnRoleLeave(r)
	})
}

func (c *MapChunk) OnRoleEnter(role *MapRole) {
	c.roles.Traversal(func(v algorithm.ElemType) {
		r := v.(*MapRole)
		role.OnRoleEnter(r)
	})
}
