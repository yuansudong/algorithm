package btree

type (
	// ListNode 表示结点元素
	ListNode struct {
		data *Element
		next *ListNode
	}
	// SortList 表示一个双向链表
	SortList struct {
		head    *ListNode
		compare func(*Element, *Element) int
		len     uint
	}
)

// CreateListNode 用于创建一个链表结点
func CreateListNode(Ele *Element) *ListNode {
	return &ListNode{
		data: Ele,
		next: nil,
	}
}

// CreateSortList 用于创建一个有序的列表
func CreateSortList(Compare func(*Element, *Element) int) *SortList {
	PList := &SortList{
		len:     0,
		compare: Compare,
		head:    CreateListNode(nil),
	}
	return PList
}

// Insert 用于向一个有序链表中插入
func (SL *SortList) Insert(InsertElement *Element) {
	PNewNode := CreateListNode(InsertElement)
	PInsertNode := SL.head.next
	PInsertNodePrev := SL.head
	for PInsertNode != nil {
		if SL.compare(PNewNode.data, PInsertNode.data) == -1 {

			break
		}
		PInsertNodePrev = PInsertNode
		PInsertNode = PInsertNode.next
	}
	PNewNode.next = PInsertNodePrev.next
	PInsertNodePrev.next = PNewNode
	SL.len = SL.len + 1
}

// Delete 用于删除指定的元素
func (SL *SortList) Delete(Ele *Element) *ListNode {
	PIterator := SL.head.next
	PIteratorPrev := SL.head
	for PIterator != nil {
		if SL.compare(PIterator.data, Ele) == 0 {
			PIteratorPrev.next = PIterator.next
			SL.len = SL.len - 1
			break
		}
		PIteratorPrev = PIterator
		PIterator = PIterator.next
	}
	return PIterator
}                         

//   GetIndex
