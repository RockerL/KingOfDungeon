package internal

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	"server/base"
)

type Map struct {
	*module.Skeleton
	Id       int32
	closeSig chan bool
}

func handleMapLoad(args []interface{}) {
	m := args[0].(*Map)
	log.Debug("handle map %v loading", m.Id)

	ChanRPC.Go("MapLoaded", m.Id)
}

func NewMap(id int32) *Map {
	m := &Map{
		Id:       id,
		Skeleton: base.NewSkeleton(),
		closeSig: make(chan bool, 0),
	}

	m.ChanRPCServer.Register("LoadMap", handleMapLoad)

	return m
}

func (m *Map) OnInit() {

}

func (m *Map) OnDestroy() {

}
