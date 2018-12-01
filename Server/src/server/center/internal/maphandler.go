package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2/bson"
	"proto"
	"shared"
	"strconv"
)

//按照大平地生成地形
func NewChunkDataWithFlat(id string, x int32, z int32) *MapChunkData {
	chunk := &MapChunkData{
		ChunkId:    id,
		ChunkX:     x,
		ChunkZ:     z,
		BlockArray: make([]BlockData, shared.ChunkBlockNum*shared.ChunkBlockNum*shared.BlockMaxY),
	}

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
				chunk.BlockArray[idx].BlockType = uint8(t)
			}
		}
	}

	return chunk
}

//处理地图初始化加载
func handleMapLoad(args []interface{}) {
	m := args[0].(*Map)
	log.Debug("handle map %v loading", m.Id)

	s := DBSession.Ref()
	defer DBSession.UnRef(s)

	mapCollectName := MapTableNamePrefix + strconv.Itoa(int(m.Id))

	for chunkX := int32(0); chunkX < shared.MaxChunkNum; chunkX++ {
		for chunkZ := int32(0); chunkZ < shared.MaxChunkNum; chunkZ++ {
			chunkName := fmt.Sprintf("%v_%v", chunkX, chunkZ)
			chunkData := new(MapChunkData)

			err := s.DB(WorldDBName).C(mapCollectName).Find(bson.M{"Id": chunkName}).One(&chunkData)
			if err != nil {
				chunkData = NewChunkDataWithFlat(chunkName, chunkX, chunkZ)
				s.DB(WorldDBName).C(mapCollectName).Insert(chunkData)
			}

			m.chunks[chunkZ*shared.MaxChunkNum+chunkX].data = chunkData
		}
	}

	ChanRPC.Go("MapLoaded", m.Id)
}

//处理角色进入地图请求
func handleRoleEnterMap(args []interface{}) {
	m := args[0].(*Map)
	roleData := args[1].(*shared.RoleData)
	agent := args[2].(gate.Agent)

	rsp := &proto.RspEnterGs{}
	roleIdx := m.roleIndex.Alloc()
	if roleIdx < 0 {
		rsp.RetCode = 4
		agent.WriteMsg(rsp)
		return
	}

	//把角色绑定到Agent
	r := &MapRole{
		m:     m,
		data:  roleData,
		idx:   roleIdx,
		agent: agent,
	}

	m.roles[roleIdx] = r
	agent.SetUserData(r)

	//返回给客户端
	rsp.RetCode = 0
	agent.WriteMsg(rsp)
}
