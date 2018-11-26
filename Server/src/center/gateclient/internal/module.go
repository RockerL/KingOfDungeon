package internal

import (
	"center/base"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
)

type Module struct {
	*module.Skeleton				//组合了Leaf的Skeleton类型，一起实现了Module
}

var (
	skeleton = base.NewSkeleton()
	ChanRPC = skeleton.ChanRPCServer
)

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	log.Debug("gate client module init")
}

func (m *Module) OnDestroy() {

}
