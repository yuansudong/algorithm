package brtree

import "reflect"

var red byte = 'r'
var black byte = 'b'

type (
	// BRTree 为红黑树的数据描述
	BRTree struct {
		empty  *treeNode
		root   *treeNode
		length int
	}
	// treeNode 为红黑树的树结点
	treeNode struct {
		color  byte
		parent *treeNode
		left   *treeNode
		right  *treeNode
		key    interface{}
		value  interface{}
	}
)

// Create 用于创建一颗红黑树
func Create() *BRTree {
	PBRTree := &BRTree{}
	PBRTree.empty = createNode(nil, nil)
	PBRTree.empty.color = black
	PBRTree.root = PBRTree.empty
	return PBRTree
}

// createNode 用于创建一个结点
func createNode(Key interface{}, value interface{}) *treeNode {
	PNode := &treeNode{}
	PNode.key = Key
	PNode.value = value
	return PNode
}

// leftRotate 用于左旋一个结点
func leftRotate(BRT *BRTree, BeNode *treeNode) {
	rightNode := BeNode.right
	BeNode.right = rightNode.left
	if rightNode.left != BRT.empty {
		rightNode.left.parent = BeNode
	}
	rightNode.parent = BeNode.parent
	// 确定rightNode 的位置
	if BeNode.parent == BRT.empty {
		// 如果满足则证明是根结点
		BRT.root = rightNode
	} else if BeNode.parent.left == BeNode {
		BeNode.parent.left = rightNode
	} else {
		BeNode.parent.right = rightNode
	}
	rightNode.left = BeNode
	BeNode.parent = rightNode
}

// rightRotate 右旋一个结点
func rightRotate(BRT *BRTree, BeNode *treeNode) {
	leftNode := BeNode.left
	BeNode.left = leftNode.right
	if leftNode.right != BRT.empty {
		leftNode.right.parent = BeNode
	}
	leftNode.parent = BeNode.parent
	// 确定BeNode 的位置
	if BeNode.parent == BRT.empty {
		//证明是父结点
		BRT.root = leftNode
	} else if BeNode.parent.left == BeNode {
		BeNode.parent.left = leftNode
	} else {
		BeNode.parent.right = leftNode
	}
	leftNode.right = BeNode
	BeNode.parent = leftNode
}

// compare 用于比较,小于返回-1,0 等于,1大于, -2 未知的比较类型
func compare(this interface{}, that interface{}) (ret int) {
	ret = -2
	thisv := reflect.ValueOf(this)
	thatv := reflect.ValueOf(that)
	if thisv.Kind() == thatv.Kind() {
		switch this.(type) {
		case int:
			if this.(int) < that.(int) {
				ret = -1
			} else if this.(int) == that.(int) {
				ret = 0
			} else {
				ret = 1
			}
		case uint:
			if this.(uint) < that.(uint) {
				ret = -1
			} else if this.(uint) == that.(uint) {
				ret = 0
			} else {
				ret = 1
			}
		case string:
			if this.(string) < that.(string) {
				ret = -1
			} else if this.(string) == that.(string) {
				ret = 0
			} else {
				ret = 1
			}
		default:
			panic("unknown type")
		}
	}
	return ret
}

// Insert 用于向红黑树中插入一个元素 O(lgn)
func Insert(BRT *BRTree, Key interface{}, Value interface{}) {
	// x表示当前结点,y表示x的父结点
	var y, x *treeNode
	y = BRT.empty
	x = BRT.root
	var compareRet int
	for x != BRT.empty {
		y = x
		compareRet = compare(Key, x.key)
		if compareRet == -1 {
			x = x.left
		} else if compareRet == 0 {
			x.value = Value
			return
		} else {
			x = x.right
		}
	}
	// 创建插入结点,并将该结点的颜色设置为红色
	InsertNode := createNode(Key, Value)
	InsertNode.color = red
	// 设置插入结点的父结点为y
	InsertNode.parent = y
	InsertNode.left = BRT.empty
	InsertNode.right = BRT.empty
	// 下面开始确定插入结点的位置
	if y == BRT.empty {
		// 证明当前红黑树中还没有结点
		BRT.root = InsertNode
	} else {
		compareRet = compare(InsertNode.key, y.key)
		if compareRet == -1 {
			y.left = InsertNode
		} else {
			y.right = InsertNode
		}
	}
	// 下面开始对整个二叉树重新着色
	fixup(BRT, InsertNode)

}

// 用于对二叉树的结点进行着色
func fixup(BRT *BRTree, InsertNode *treeNode) {
	/*
	   (1) 如果InsertNode 是根结点,那么其父结点的颜色是黑色,所以不满足

	*/
	for InsertNode.parent.color == red {
		// 首先判断父结点是左结点还是右结点
		if InsertNode.parent == InsertNode.parent.parent.left {
			// 获得插入结点的父结点的兄弟结点
			y := InsertNode.parent.parent.right
			// 第一种情况 父结点的兄弟结点是红色
			if y.color == red {
				// 如果满足则需要做下面的事情
				// 将插入结点的父结点设置为黑色
				InsertNode.parent.color = black
				// 将插入节点的父结点的兄弟结点设置为
				y.color = black
				InsertNode.parent.parent.color = red
				InsertNode = InsertNode.parent.parent
			} else {
				if InsertNode == InsertNode.parent.right {
					// 第二种情况 插入结点的叔伯结点是黑色,且本身是一个右结点
					InsertNode = InsertNode.parent
					leftRotate(BRT, InsertNode)
				}
				// 第三种情况 插入结点的叔伯结点是黑色, 且本身是一个左结点
				InsertNode.parent.color = black
				InsertNode.parent.parent.color = red
				rightRotate(BRT, InsertNode.parent.parent)
			}
		} else {
			y := InsertNode.parent.parent.left
			// 第一种情况 父结点的兄弟结点是红色
			if y.color == red {
				// 如果满足则需要做下面的事情
				// 将插入结点的父结点设置为黑色
				InsertNode.parent.color = black
				// 将插入节点的父结点的兄弟结点设置为
				y.color = black
				InsertNode.parent.parent.color = red
				InsertNode = InsertNode.parent.parent
			} else {
				if InsertNode == InsertNode.parent.left {
					// 第二种情况 插入结点的叔伯结点是黑色,且本身是一个右结点
					InsertNode = InsertNode.parent
					rightRotate(BRT, InsertNode)
				}
				// 第三种情况 插入结点的叔伯结点是黑色, 且本身是一个左结点
				InsertNode.parent.color = black
				InsertNode.parent.parent.color = red
				leftRotate(BRT, InsertNode.parent.parent)
			}

		}
	}
	BRT.root.color = black
}

// replace 用于替换红黑树中的结点,BRT为红黑树结构,而U为被替换的结点,v为替换结点
func replace(BRT *BRTree, U, V *treeNode) {
	if U.parent == BRT.empty {
		// 证明要替换的结点是根结点
		BRT.root = V
	} else if U == U.parent.left {
		// 证明要替换的结点是一个左结点
		U.parent.left = V
	} else {
		U.parent.right = V
	}
	// 替换父结点
	V.parent = U.parent
}

// search 用于搜索指定的key的结点
func search(BRT *BRTree, Key interface{}) *treeNode {
	ret := BRT.empty
	var compareRet int
	for x := BRT.root; x != BRT.empty; {
		compareRet = compare(Key, x.key)
		if compareRet == -1 {
			// 证明key 是小于 x.key
			x = x.left
		} else if compareRet == 0 {
			// 证明key 是等于 x.key 的
			return x
		} else {
			// 证明key 是大小与 x.key
			x = x.right
		}
	}
	return ret
}

// getMiniNode 获得指定结点的最小的元素
func getMiniNode(BRT *BRTree, Node *treeNode) *treeNode {

	var iNode *treeNode
	for iNode = Node; iNode.left != BRT.empty; {

	}
	return iNode

}

// getMaxNode 获得指定结点的最大的元素
func getMaxNode(BRT *BRTree, Node *treeNode) *treeNode {

	var iNode *treeNode
	for iNode = Node; iNode.right != BRT.empty; iNode = iNode.right {

	}
	return iNode
}

// Delete 用于删除指定的key,x 为着色点, y为后继点
func Delete(BRT *BRTree, Key interface{}) {
	if DeleteNode := search(BRT, Key); DeleteNode != BRT.empty {
		y := DeleteNode
		x := BRT.empty
		yOriColor := y.color
		if DeleteNode.left == BRT.empty {
			x = DeleteNode.right
			replace(BRT, DeleteNode, x)
		} else if DeleteNode.right == BRT.empty {
			x = DeleteNode.left
			replace(BRT, DeleteNode, x)
		} else {
			y = getMiniNode(BRT, DeleteNode.right)
			yOriColor = y.color
			x = y.right
			if y.parent == DeleteNode {
				x.parent = y
			} else {
				replace(BRT, y, y.right)
				y.right = DeleteNode.right
				y.right.parent = y
			}
			replace(BRT, DeleteNode, y)
			y.left = DeleteNode.left
			y.left.parent = y
			y.color = DeleteNode.color
		}
		if yOriColor == black {
			deleteFixup(BRT, x)
		}
	}

}

// Get 用于获取指定的key的value
func Get(BRT *BRTree, Key interface{}) interface{} {
	return search(BRT, Key).value
}

// deleteFixup 用于进行着色
func deleteFixup(BRT *BRTree, x *treeNode) {
	for x != BRT.root && x.color == black {
		if x == x.parent.left {
			// w 为x的兄弟结点
			w := x.parent.right
			if w.color == red {
				w.color = black
				w.parent.color = red
				leftRotate(BRT, x.parent)
				w = x.parent.right
			}
			if w.left.color == black && w.right.color == black {
				w.color = red
				x = x.parent
			} else {
				if w.right.color == black {
					w.left.color = black
					w.color = red
					rightRotate(BRT, w)
					w = x.parent.right
				}
				w.color = x.parent.color
				x.parent.color = black
				w.right.color = black
				leftRotate(BRT, x.parent)
			}
		} else {
			w := x.parent.left
			if w.color == red {
				w.color = black
				w.parent.color = red
				rightRotate(BRT, x.parent)
				w = x.parent.left
			}
			if w.right.color == black && w.left.color == black {
				w.color = red
				x = x.parent
			} else {
				if w.left.color == black {
					w.right.color = black
					w.color = red
					leftRotate(BRT, w)
					w = x.parent.left
				}
				w.color = x.parent.color
				x.parent.color = black
				w.right.color = black
				rightRotate(BRT, x.parent)
			}
		}
		x = BRT.root
	}

}
