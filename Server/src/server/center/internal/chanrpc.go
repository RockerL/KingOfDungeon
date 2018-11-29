package internal

import (
	"github.com/name5566/leaf/log"
	"proto"
	"server/conf"
)

func init() {
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC("MapLoaded", rpcMapLoaded)
}

//当和center server断开
func rpcCloseAgent(args []interface{}) {
	a := args[0].(*TCPAgent)
	_ = a

	log.Release("disconnect from center server")
}

//当某个地图加载完毕后通知
func rpcMapLoaded(args []interface{}) {
	mapId := args[0].(int32)
	log.Debug("map loaded %v", mapId)

	loadedMapCount++
	if loadedMapCount == len(maps) {

		log.Debug("all map loaded, count = %v", loadedMapCount)

		ctAgent.WriteMsg(&proto.NotifyServerInited{
			LoadedMaps: conf.Server.MapLoad,
		})
	}
}
