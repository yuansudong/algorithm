package radixTree

// RadixTree 用于描述一个基数树
type RadixTree struct {
	root *radixTreeNode
	// high 用于表示高度
	high   int
	length int
}

//

// Create 用于创建一个Radixtree
func Create(H int) *RadixTree {
	PRT := &RadixTree{
		root:   createRadixTreeNode(nil),
		length: 0,
		high:   H,
	}
	createSubTree(1, PRT.high, PRT.root)
	return PRT
}
func createSubTree(curHigh int, totalHigh int, parent *radixTreeNode) {
	if curHigh > totalHigh {

	} else {
		for i := 0; i < charsetLen; i++ {
			parent.benchMark[i] = createRadixTreeNode(nil)
			createSubTree(curHigh+1, totalHigh, parent.benchMark[i])
		}
	}

}

// Put 用于向RaidxTree中放置一个key-val 结构体,如果该key存在则覆盖,如果不存在则添加
func (rt *RadixTree) Put(key string, value interface{}) bool {
	if radixTreeNode

	getRadixTreeNode(0, len(key), key, rt.root)
}

func getRadixTreeNode(Index int, Length int, Key string, Node *radixTreeNode) *radixTreeNode {
	if Index == Length {
		return Node
	}
	SubKey := Key[Index]
	SubKeyIndex := getCharsetIndex(SubKey)
	if SubKeyIndex == -1 {
		return nil
	}
	return getRadixTreeNode(Index+1, Length, Key, Node.benchMark[SubKeyIndex])
}

// radixTreeNode 用于描述一个radix 树节点
type radixTreeNode struct {
	// 数据
	data interface{}
	// 分支
	benchMark []*radixTreeNode
}

// 该函数用于创建基数树节点
func createRadixTreeNode(D interface{}) *radixTreeNode {
	PTreeNode := &radixTreeNode{data: D}
	PTreeNode.benchMark = make([]*radixTreeNode, charsetLen)
	return PTreeNode
}
