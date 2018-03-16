package ttree

import (
	"reflect"
)

// InOrder 中序遍历
func InOrder(T *TTree, set []interface{}, call func(key interface{}, value interface{}, set []interface{}) []interface{}) (result []interface{}) {
	if T.root != nil {
		result = inOrder(T.root, set, call)
	}
	return result
}
func inOrder(Node *tTreeNode, set []interface{}, call func(key interface{}, value interface{}, set []interface{}) []interface{}) []interface{} {
	if Node != nil {
		set = inOrder(Node.left, set, call)
		set = call(Node.key, Node.value, set)
		set = inOrder(Node.right, set, call)
	}
	return set
}

// PreOrder 先序遍历
func PreOrder(T *TTree, set []interface{}, call func(key interface{}, value interface{}, set []interface{}) []interface{}) (result []interface{}) {
	if T.root != nil {
		result = preOrder(T.root, set, call)
	}
	return result
}
func preOrder(Node *tTreeNode, set []interface{}, call func(key interface{}, value interface{}, set []interface{}) []interface{}) []interface{} {
	if Node != nil {
		set = inOrder(Node.left, set, call)
		set = inOrder(Node.right, set, call)
		set = call(Node.key, Node.value, set)
	}
	return set
}

// PostOrder 后续遍历
func PostOrder(T *TTree, set []interface{}, call func(key interface{}, value interface{}, set []interface{}) []interface{}) (result []interface{}) {
	if T.root != nil {
		result = postOrder(T.root, set, call)
	}
	return result
}
func postOrder(Node *tTreeNode, set []interface{}, call func(key interface{}, value interface{}, set []interface{}) []interface{}) []interface{} {
	if Node != nil {
		set = inOrder(Node.left, set, call)
		set = inOrder(Node.right, set, call)
		set = call(Node.key, Node.value, set)
	}
	return set
}

// get 用于查找key对应的结点, 如果有则返回的是值,否则返回的是nil的interface
func getNode(T *TTree, key interface{}) *tTreeNode {
	var ele *tTreeNode
	for ele = T.root; ele != nil && ele.key != key; {
		if compare(key, ele.key) == -1 {
			ele = ele.left
		} else {
			ele = ele.right
		}
	}
	return ele
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

// getMiniElement 用于获取key值最小的元素,若树为空则返回的key以及value都是nil
func getMiniNode(T *tTreeNode) *tTreeNode {

	if T != nil {
		var X *tTreeNode
		for X = T; X.left != nil; X = X.left {
			// TODO: 什么都不做
		}
		return X
	}
	return nil
}

// GetMaxElement 用于获取key值最大的元素,若树为空则返回的key以及value都是nil
func getMaxElement(T *TTree) *tTreeNode {
	if T.root != nil {
		var X *tTreeNode
		for X = T.root; X.right != nil; X = X.right {
			// TODO: 什么都不做
		}
		return X
	}
	return nil
}

// getSuccessor 获得继承者
func getSuccessor(x *tTreeNode) *tTreeNode {
	if x != nil {
		if x.right != nil {
			return getMiniNode(x)
		}
		y := x.parent
		for y != nil && x == y.right {
			x = y
			y = y.parent
		}
		return y
	}
	return nil
}

// Update 用于向二叉树中更新数据, 不存在插入,存在则修改
func Update(T *TTree, key interface{}, value interface{}) {
	var x, y *tTreeNode
	x = T.root
	for x != nil {
		y = x
		ret := compare(key, x.key)
		if ret == -1 {
			x = x.left
		} else if ret == 0 {
			x.value = value
			return
		} else {
			x = x.right
		}
	}
	PNode := createTTreeNode()
	PNode.key = key
	PNode.value = value
	PNode.parent = y
	if y == nil {
		T.root = PNode
	} else if compare(PNode.key, y.key) == -1 {
		y.left = PNode
	} else {
		y.right = PNode
	}
}

/*
   从一颗二叉搜索树T中删除指定的元素分为三种情况
   1. 如果要删除的元素中没有孩子结点,那么只是将其删除
	  用nil来替换该元素
   2. 如果要删除的元素只有一个孩子,那么让这个孩子替代被删除
      元素的位置,并修改这个孩子的父结点
   3. 如果z有两个孩子,那么需要找到z的后继(一定在被删除元素的右子树中)

*/

// transparent 用于替换将U结点替换为V结点,替换过程中,需要
// 修改U的父结点的指向,以及V结点的父结点的存储
func transparent(T *TTree, U *tTreeNode, V *tTreeNode) {

	if U.parent == nil {
		// 如果满足条件,那么代表U结点是一个根节点
		T.root = V
	} else if U == U.parent.left {
		// 如果条件满足,那么代表U结点是一个左结点
		U.parent.left = V
	} else {
		// 如果条件满足,那么代表这个结点是右结点
		U.parent.right = V
	}
	// 设置V的父结点
	if V != nil {
		V.parent = U.parent
	}

}

// Delete 时间复杂度为O(h) 用于删除指定key值的, 如果有key就删除,无key就不删除
func Delete(T *TTree, key interface{}) {
	if DeleteNode := getNode(T, key); DeleteNode != nil {
		if DeleteNode.left == nil {
			// 如果只有右结点
			transparent(T, DeleteNode, DeleteNode.right)
		} else if DeleteNode.right == nil {
			// 如果只有左结点
			transparent(T, DeleteNode, DeleteNode.left)
		} else {
			// 如果有两个结点,则先找到被删除结点中的后继结点,该
			// 后继结点,一定在被删除结点的右节点的最左层结点
			AfterNode := getMiniNode(DeleteNode.right)
			if AfterNode.parent != DeleteNode {
				// 如果后继点的父结点不是删除结点
				// 先将后继结点替换为后继结点的右结点
				transparent(T, AfterNode, AfterNode.right)
				AfterNode.right = DeleteNode.right
				AfterNode.right.parent = AfterNode
			}
			transparent(T, DeleteNode, AfterNode)
			AfterNode.left = DeleteNode.left
			AfterNode.left.parent = AfterNode
		}
	}

}
