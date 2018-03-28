package btree

// SortList 有序列表
type SortList struct {
	len  int
	head *SortListNode
}

// CreateSortList 用于创建一个有序的列表
func CreateSortList() *SortList {
	PSortList := new(SortList)
	PSortList.len = 0
	PSortList.head = new(SortListNode)
	return PSortList
}

// SortListNode 用于表示list的结点
type SortListNode struct {
	key        BKey
	value      interface{}
	prev, next *SortListNode
}

// Update 用于向有序列表中增加或者更新一个元素,且Key互斥
func (L *SortList) Update(Key *BKey, Value interface{}) {
	PHead := L.head.next
	PPrev := L.head
	var CompareRet int
	for PHead != nil {
		CompareRet = PHead.key.compare(Key)
		if CompareRet == 1 {
			break
		} else if CompareRet == 0 {
			PHead.value = Value
			return
		}
		PPrev = PHead
		PHead = PHead.next
	}
	PNewSortListNode := new(SortListNode)
	PNewSortListNode.key = *Key
	PNewSortListNode.value = Value
	// 如果满足条件则代表是在尾部结点插入
	PPrev.next = PNewSortListNode
	PNewSortListNode.next = PHead
	PNewSortListNode.prev = PPrev
	if PHead != nil {
		PHead.prev = PNewSortListNode
	}
	L.len++

}

// Delete 用于删除指定的Key
func (L *SortList) Delete(Key *BKey) {
	PHead := L.head.next
	PPrev := L.head
	var CompareRet int
	for PHead != nil {
		CompareRet = PHead.key.compare(Key)
		if CompareRet == 0 {
			break
		}
		PPrev = PHead
		PHead = PHead.next
	}
	// 如果满足条件则代表是在尾部结点插入
	if PHead != nil {
		PPrev.next = PHead.next
		if PHead.next != nil {
			PHead.next.prev = PPrev
		}
		L.len--
	}
}
