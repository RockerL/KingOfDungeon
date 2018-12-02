package algorithm

type IndexAllocator struct {
	useList  *DoubleList
	freeList *DoubleList
}

//新建一个索引分配器，传入为索引值最大值+1
func NewIndexAllocator(maxSize int) *IndexAllocator {
	if maxSize <= 0 {
		return nil
	}

	c := &IndexAllocator{
		useList:  NewLinkList(maxSize),
		freeList: NewLinkList(maxSize),
	}

	//初始空闲列表里有所有索引值
	for i := 0; i < maxSize; i++ {
		c.freeList.Add(i)
	}

	return c
}

//申请一个索引
func (c *IndexAllocator) Alloc() int {
	if c.freeList.size <= 0 {
		return -1
	}

	idx := c.freeList.GetFirst()
	c.freeList.RemoveFirst()
	c.useList.Add(idx)

	return idx
}

//释放一个索引
func (c *IndexAllocator) Free(idx int) {
	if c.useList.Remove(idx) {
		c.freeList.Add(idx)
	}
}
