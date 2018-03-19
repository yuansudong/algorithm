package intervaltree

var red byte = 'r'
var black byte = 'b'

type (
	// IntervalTree 为红黑树的数据描述
	IntervalTree struct {
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
		// 区间
		interval *Range
		// 以该结点为根结点的元素最多能存储的最大值
		max int
	}
)

// intervalSearch 用于搜索指定的区间
func intervalSearch(IT *IntervalTree, i *Range) *treeNode {
	x := IT.root
	for x != IT.empty && !x.interval.Overlap(i) {
		if x.left != IT.empty && i.GetLow() <= x.left.max {
			x = x.left
		} else {
			x = x.right
		}
	}
	return x
}

// intervalInsert 用于向区间树中插入一个区间
func intervalInsert(IT *IntervalTree, low, high int) {

}
