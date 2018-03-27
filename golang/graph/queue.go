package graph

type (
	// Queue 用于表示一个队列
	Queue struct {
		len  int
		head *Element
		tail *Element
	}
	// Element 用于表示队列中的一个元素
	Element struct {
		value interface{}
		next  *Element
	}
)

// CreateElement 用于创建一个元素
func CreateElement(V interface{}) *Element {
	return &Element{
		value: V,
	}
}

// CreateQueue 用于创建一个Queue
func CreateQueue() *Queue {
	PQueue := &Queue{}
	PQueue.head = CreateElement(nil)
	PQueue.tail = PQueue.head
	PQueue.len = 0
	return PQueue
}

// EnQueue 用于向队列中添加一个元素
func (Q *Queue) EnQueue(V interface{}) {
	PNewElement := CreateElement(V)
	Q.tail.next = PNewElement
	Q.tail = PNewElement
	Q.len++
}

// DeQueue 用于向队列中删除一个元素并且返回
func (Q *Queue) DeQueue() (bool, interface{}) {
	if Q.len != 0 {
		PEle := Q.head.next
		Q.head.next = Q.head.next.next
		Q.len--
		return true, PEle.value

	}
	return false, nil
}
