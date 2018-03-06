package skiplist

type skipListNodeData struct {
	key   interface{} // key 应该是唯一的
	value interface{}
	score float64
}

// equal 用于比较两个结构体数据是不是相等的
func (slnd *skipListNodeData) equal(other *skipListNodeData) bool {

	return slnd.key == other.key
}

// 创建一个结点数据
func createSkipListNodeData(Score float64, Key interface{}, Value interface{}) *skipListNodeData {
	return &skipListNodeData{
		key:   Key,
		value: Value,
		score: Score,
	}
}

type skipListNode struct {
	obj   *skipListNodeData
	up    *skipListNode
	down  *skipListNode
	left  *skipListNode
	right *skipListNode
}

//
func makeNode(Num int, Object *skipListNodeData) []*skipListNode {
	skipListNodes := make([]skipListNode, Num)
	skipListNodePointers := make([]*skipListNode, Num)
	range1 := Num - 1
	for i := 1; i < range1; i++ {
		skipListNodePointers[i] = &skipListNodes[i]
		skipListNodePointers[i].obj = Object
		skipListNodePointers[i].up = skipListNodePointers[i+1]
		skipListNodePointers[i].down = skipListNodePointers[i-1]
	}
	skipListNodePointers[0] = &skipListNodes[0]
	skipListNodePointers[0].obj = Object
	if range1 != 0 {
		skipListNodePointers[range1] = &skipListNodes[range1]
		skipListNodePointers[range1].obj = Object
		skipListNodePointers[0].up = skipListNodePointers[1]
		skipListNodePointers[0].down = nil
		skipListNodePointers[range1].up = nil
		skipListNodePointers[range1].down = skipListNodePointers[range1-1]
	}

	return skipListNodePointers

}

//  创建一个跳跃链表的结点
func createSkipListNode(Object *skipListNodeData) *skipListNode {
	return &skipListNode{
		obj:   Object,
		up:    nil,
		down:  nil,
		left:  nil,
		right: nil,
	}

}
