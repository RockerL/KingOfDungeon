package algorithm

type LinkList struct {
	node
	dummyHead *node  //虚拟头节点
	maxSize   int    //最大节点个数
	array     []node //节点容器
	size      int    //有效节点个数
}

type node struct {
	Index int
	Next  *node
	Prev  *node
}

func NewLinkList(maxSize int) *LinkList {
	l := &LinkList{
		dummyHead: &node{-1, nil, nil},
		array:     make([]node, maxSize),
		maxSize:   maxSize,
		size:      0,
	}

	for i := 0; i < l.maxSize; i++ {
		l.array[i].Index = i
	}

	return l
}

func (link *LinkList) MaxSize() int {
	return link.maxSize
}

func (link *LinkList) Size() int {
	return link.size
}

//根据值来移除节点
func (link *LinkList) Remove(index int) bool {
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

func (link *LinkList) RemoveFirst() {
	if link.dummyHead.Next == nil {
		return
	}

	link.Remove(link.dummyHead.Next.Index)
}

func (link *LinkList) Add(index int) {
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

func (link *LinkList) GetFirst() int {
	if link.dummyHead.Next == nil {
		return link.dummyHead.Index
	}

	return link.dummyHead.Next.Index
}
