package internal

import (
	"server/conf"
	"shared"
)

func (c *MapChunk) InitChunkWithTerrain(id string, chunkX int32, chunkZ int32) {
	c.data.ChunkId = id
	c.data.ChunkX = chunkX
	c.data.ChunkZ = chunkZ
	c.blockStartX = chunkX * shared.ChunkBlockNum
	c.blockStartZ = chunkZ * shared.ChunkBlockNum

	//生成地表高度图
	heightMap := make([]uint8, shared.ChunkBlockNum*shared.ChunkBlockNum)
	for z := int32(0); z < shared.ChunkBlockNum; z++ {
		for x := int32(0); x < shared.ChunkBlockNum; x++ {
			idx := z*shared.ChunkBlockNum + x
			noise := (0.5 + c.m.perlinNoise.Noise2D(float64(c.blockStartX+x), float64(c.blockStartZ+z))) * shared.BlockMaxY
			heightMap[idx] = uint8(noise)
		}
	}

	//填充地表以下的块为土块，高于地表的为空气块
	for y := uint8(0); y < shared.BlockMaxY; y++ {
		for z := 0; z < shared.ChunkBlockNum; z++ {
			for x := 0; x < shared.ChunkBlockNum; x++ {
				var t uint16 = BlockAir
				heightIdx := z*shared.ChunkBlockNum + x
				if y <= heightMap[heightIdx] {
					t = BlockEarth
				}

				idx := int(y)*shared.ChunkBlockNum*shared.ChunkBlockNum + z*shared.ChunkBlockNum + x
				c.data.BlockArray[idx].BlockType = t
				c.data.BlockArray[idx].Content = conf.GetContentInitValue(t)
			}
		}
	}
}
