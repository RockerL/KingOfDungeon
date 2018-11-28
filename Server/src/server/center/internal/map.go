package internal

import (
	"github.com/name5566/leaf/module"
	"server/base"
)

type Map struct {
	Id int
	*module.Skeleton
}

func NewMap(id int) *Map {
	return &Map{
		Id:       id,
		Skeleton: base.NewSkeleton(),
	}
}

func (m *Map) OnInit() {

}

func (m *Map) OnDestroy() {

}
