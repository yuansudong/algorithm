package skiplist

// SkipList 用于描述一个跳表
type SkipList struct {
	header []*skipListNode
	tail   []*skipListNode
	length int
	level  int
}

// Create 用于创建一个跳表
func Create() *SkipList {
	DefaultObject := createSkipListNodeData(0, nil, nil)
	PSkipList := &SkipList{}
	PSkipList.level = 0
	PSkipList.length = 0
	PSkipList.header = makeNode(skipListMaxLevel, DefaultObject)
	PSkipList.tail = makeNode(skipListMaxLevel, DefaultObject)
	for i := 1; i < skipListMaxLevel; i++ {
		PSkipList.header[i].right = PSkipList.tail[i]
		PSkipList.tail[i].left = PSkipList.header[i]
	}
	return PSkipList
}

// InsertFromFront 用于从头部插入元素
func (sl *SkipList) InsertFromFront(Object *skipListNodeData) {
	// 先随机出层数, 该层数代表着从哪一层开始构建元素,levelInt最大为31
	levelInt := randomLevel()
	InsertNodes := makeNode(levelInt+1, Object)
	// 构建插入元素
	var Tail *skipListNode
	var Head *skipListNode
	var InsertNode *skipListNode
	var Next *skipListNode
	StartNode := sl.header[levelInt]
	for i := levelInt; i >= 0; i-- {
		Tail = sl.tail[i]
		Head = sl.header[i]
		InsertNode = InsertNodes[levelInt]
		for {
			if StartNode.right == Tail {
				break
			}
			// 找出第一个大于Object分值的元素
			if StartNode != Head && Object.score > StartNode.obj.score && Object.score <= StartNode.right.obj.score {
				break
			}
			StartNode = StartNode.right
		}
		Next = StartNode.right
		StartNode.right = InsertNode
		InsertNode.left = StartNode
		Next.left = InsertNode
		InsertNode.right = Next
		StartNode = StartNode.down
	}
}

// InsertFromTail 用于从尾部插入元素
func (sl *SkipList) InsertFromTail(Object *skipListNodeData) {
	// 先随机出层数, 该层数代表着从哪一层开始构建元素,levelInt最大为31
	levelInt := randomLevel()
	InsertNodes := makeNode(levelInt+1, Object)
	// 构建插入元素
	var Tail *skipListNode
	var Head *skipListNode
	var InsertNode *skipListNode
	var Prev *skipListNode
	StartNode := sl.header[levelInt]
	for i := levelInt; i >= 0; i-- {
		Tail = sl.tail[i]
		Head = sl.header[i]
		InsertNode = InsertNodes[levelInt]
		for {
			if StartNode.left == Head {
				break
			}
			// 找出第一个大于Object分值的元素
			if StartNode != Tail && Object.score < StartNode.obj.score && Object.score >= StartNode.right.obj.score {
				break
			}
			StartNode = StartNode.left
		}
		Prev = StartNode.left

		InsertNode.right = StartNode
		StartNode.left = InsertNode
		InsertNode.left = Prev
		Prev.right = InsertNode
		StartNode = StartNode.down
	}
}
