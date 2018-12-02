package algorithm

type DoubleList struct {
	dummyHead *DNode  //虚拟头节点
	maxSize   int     //最大节点个数
	array     []DNode //节点容器
	size      int     //有效节点个数
}

type DNode struct {
	Index int
	Next  *DNode
	Prev  *DNode
}

func NewLinkList(maxSize int) *DoubleList {
	l := &DoubleList{
		dummyHead: &DNode{-1, nil, nil},
		array:     make([]DNode, maxSize),
		maxSize:   maxSize,
		size:      0,
	}

	for i := 0; i < l.maxSize; i++ {
		l.array[i].Index = i
	}

	return l
}

func (link *DoubleList) MaxSize() int {
	return link.maxSize
}

func (link *DoubleList) Size() int {
	return link.size
}

//根据值来移除节点
func (link *DoubleList) Remove(index int) bool {
	if index < 0 || index > link.maxSize {
		panic("remove out of range")
		return false
	}

	n := link.array[index]

	if n.Index != index {
		panic("remove index error")
		return false
	}

	if n.Prev != nil {
		n.Prev.Next = n.Next
	}
	if n.Next != nil {
		n.Next.Prev = n.Prev
	}

	n.Next = nil
	n.Prev = nil

	link.size--

	return true
}

func (link *DoubleList) RemoveFirst() {
	if link.dummyHead.Next == nil {
		return
	}

	link.Remove(link.dummyHead.Next.Index)
}

func (link *DoubleList) Add(index int) {
	if index < 0 || index > link.maxSize {
		panic("add out of range")
		return
	}

	n := link.array[index]
	if n.Next != nil || n.Prev != nil {
		panic("add duplicate node")
		return
	}

	n.Next = link.dummyHead.Next
	n.Prev = link.dummyHead
	link.dummyHead.Next = &n

	link.size++
}

func (link *DoubleList) GetFirst() int {
	if link.dummyHead.Next == nil {
		return link.dummyHead.Index
	}

	return link.dummyHead.Next.Index
}
