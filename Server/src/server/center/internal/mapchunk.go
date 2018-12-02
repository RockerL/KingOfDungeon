package internal

import "proto"

//运行时地图块
type MapChunk struct {
	data *MapChunkData //地图区块数据
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
