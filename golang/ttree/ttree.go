package ttree

type tTreeNode struct {
	key    interface{}
	value  interface{}
	left   *tTreeNode
	right  *tTreeNode
	parent *tTreeNode
}

func createTTreeNode() *tTreeNode {
	return &tTreeNode{}
}

// TTree 用于描述一个二叉树的结构体
type TTree struct {
	root *tTreeNode
}

// Create 用于创建一颗二叉树
func Create() *TTree {
	return &TTree{}
}
