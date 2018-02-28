package list

// LinkedList is defined struct for linked lists
type LinkedList struct {
	size uint64
	head *LinkedElement
	tail *LinkedElement
}

// LinkedInit 用于单向链表的初始化
func (L *LinkedList) LinkedInit() {
	L.head = nil
	L.tail = nil
	L.size = 0
}

// LinkedDestory 用于单向链表的销毁
func (L *LinkedList) LinkedDestory() {
	for i := uint64(0); i < L.LinkedSize(); i++ {
		BeRemElement := L.head
		L.head = L.head.next
		BeRemElement = nil
	}
}

func (L *LinkedList) LinkedSize() uint64 {
	return L.size
}
