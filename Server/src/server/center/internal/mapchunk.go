package internal

import (
	"proto"
	"server/conf"
	"shared"
	"shared/algorithm"
)

type MapObject struct {
}

//运行时地图块
type MapChunk struct {
	data        MapChunkData //地图区块数据
	blockStartX int32
	blockStartZ int32
	roles       *algorithm.SingleList //角色的链表头
}

//按照大平地生成地形
func (c *MapChunk) InitChunkWithFlat(id string, chunkX int32, chunkZ int32) {
	c.data.ChunkId = id
	c.data.ChunkX = chunkX
	c.data.ChunkZ = chunkZ
	c.blockStartX = chunkX * shared.ChunkBlockNum
	c.blockStartZ = chunkZ * shared.ChunkBlockNum
	for y := 0; y < shared.BlockMaxY; y++ {
		for z := 0; z < shared.ChunkBlockNum; z++ {
			for x := 0; x < shared.ChunkBlockNum; x++ {
				var t uint16 = Air
				if y < shared.BlockMaxY/2 {
					t = Earth
				}
				idx := y*shared.ChunkBlockNum*shared.ChunkBlockNum + z*shared.ChunkBlockNum + x
				c.data.BlockArray[idx].BlockType = t
				c.data.BlockArray[idx].Content = conf.GetContentInitValue(t)
			}
		}
	}
}

func (c *MapChunk) MakeChunkBlockInfo() []*proto.BlockInfo {
	info := make([]*proto.BlockInfo, len(c.data.BlockArray))

	for i := 0; i < len(c.data.BlockArray); i++ {
		info[i] = &proto.BlockInfo{
			BlockType: uint32(c.data.BlockArray[i].BlockType),
			Content:   c.data.BlockArray[i].Content,
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

func (c *MapChunk) GetBlock(blockX int32, blockY int32, blockZ int32) *BlockData {
	if blockX < c.blockStartX || blockZ < c.blockStartZ ||
		blockX >= c.blockStartX+shared.ChunkBlockNum ||
		blockZ >= c.blockStartZ+shared.ChunkBlockNum ||
		blockY < 0 || blockY >= shared.BlockMaxY {
		return nil
	}

	blockX = blockX - c.blockStartX
	blockZ = blockZ - c.blockStartZ
	idx := blockY*shared.ChunkBlockNum*shared.ChunkBlockNum + blockZ*shared.ChunkBlockNum + blockX
	return &c.data.BlockArray[idx]
}
