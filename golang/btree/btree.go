package btree

type (
	// BTree 用于描述一个BTree的结构体
	BTree struct {
		M, T, Len int
		root      *TreeNode
	}
	// TreeNode 用于描述一个BTreeNode的结构
	TreeNode struct {
		keys   []datanode
		parent *TreeNode
		childs []*TreeNode
		isLeaf bool
	}
)

// CreateBTree 用于创建一颗B树
func CreateBTree(T int) *BTree {
	PBtree := new(BTree)
	PBtree.T = T
	PBtree.M = 2*T - 1
	PBtree.root = new(TreeNode)
	return PBtree
}

//createTreeNode 用于创建一个树结点
func createTreeNode(BT *BTree) *TreeNode {

}

// search 用于搜索指定根节点下的Key
func (B *BTree) search(TN *TreeNode, Key string) {
	PSearchNode := new(SortListNode)
	PHead := TN.KeyArray.head.next
	Index := 0
	var ComRet int
	for PHead != nil {
		ComRet = PSearchNode.Compare(PHead)
		if ComRet != 1 {
			break
		}
		Index++
		PHead = PHead.next
	}
	if PHead != nil && ComRet == 0 {

	}

}
