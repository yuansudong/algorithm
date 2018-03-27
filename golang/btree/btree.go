package btree

type (
	// BTree 用于描述一个B树的结构体
	BTree struct {
		m, t    int
		compare func(*Element, *Element) int
		root    *treeNode
	}
	// treeNode 用于一个树形结构结点
	treeNode struct {
		keys          *SortList
		isLeaf        bool
		childNode     []*treeNode
		parent, child *treeNode
	}
)

func createTreeNode(m int, compare func(this, that *Element) int) *treeNode {
	return &treeNode{
		keys:      CreateSortList(compare),
		isLeaf:    false,
		childNode: make([]*treeNode, m+1),
	}
}

// CreateBTree 用于创建一颗B树
func CreateBTree(t, m int, compare func(this, that *Element) int) *BTree {
	PBTree := &BTree{
		t:       m / 2,
		m:       m,
		compare: compare,
		root:    createTreeNode(m, compare),
	}
	PBTree.root.isLeaf = true
	return PBTree
}

// Insert 用于向BTree中插入一个元素
func (B *BTree) Insert(Key, Value interface{}) {
	PNewElement := CreateElement(Key, Value)
	root := B.root
	if root.isLeaf {
		// 进入这里,代表着此时root是一个叶子结点
		root.keys.Insert(PNewElement)
		//treeNodeSplit(B, root)
	} else {

	}

}
