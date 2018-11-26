package internal

import (
	"center/base"
	"github.com/name5566/leaf/module"
)

var (
	skeleton   = base.NewSkeleton()
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton

}

func (m *Module) OnDestroy() {

}
