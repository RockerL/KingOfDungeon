package algorithm

type ElemType interface{}
type ListTraversalFunc func(v ElemType)

//结点
type SNode struct {
	Data ElemType
	Next *SNode
}

//链表
type SingleList struct {
	Head *SNode //头结点
}

func NewSingleList() *SingleList {
	head := new(SNode)
	return &SingleList{head}
}

//遍历列表
func (list *SingleList) Traversal(f ListTraversalFunc) {
	p := list.Head.Next //第一个结点
	for p != nil {
		f(p.Data)
		p = p.Next
	}
}

//头部插入元素
func (list *SingleList) Insert(v ElemType) {
	p := list.Head
	s := &SNode{v, p.Next}
	p.Next = s
}

//删除结点
func (list *SingleList) Delete(v ElemType) bool {
	p := list.Head
	for p.Next != nil {
		if p.Next.Data == v {
			p.Next = p.Next.Next
			return true
		}
		p = p.Next
	}
	return false
}

//判断链表是否为空
func (list *SingleList) IsEmpty() bool {
	return list.Head.Next == nil
}

//链表长度
func (list *SingleList) Len() int {
	length := 0
	p := list.Head.Next
	for p != nil {
		p = p.Next
		length++
	}

	return length
}
