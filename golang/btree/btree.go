package btree

const (
	// BTreeSearchMinSize 最小搜索大小
	BTreeSearchMinSize = 16
	// BTreeIterMaxDepth 最大搜素深度
	BTreeIterMaxDepth = 64
)

type (
	bnode struct {
		isLeaf bool
		size   int        // 节点包含的元素数目
		eles   []*Element // 元素
		sub    []*bnode   // 下级子树
		parent *bnode
	}

	// BTree 用于B树的数据结构
	BTree struct {
		minDegree, minElementCount, maxElementCount, maxSubCount int    // BTree 的最小度数,以及最大元素个数
		root                                                     *bnode // BTree的根结点
	}
)

// nodeCreate 用于创建一个树的结点
func (B *BTree) nodeCreate(IsLeaf bool) *bnode {
	PNode := &bnode{
		isLeaf: IsLeaf,
		size:   0,
		eles:   make([]*Element, B.maxElementCount),
		sub:    make([]*bnode, B.maxSubCount),
	}
	return PNode
}

// Create 用于创建一颗B树
func Create(MinDegree int) *BTree {
	PBtree := &BTree{
		minDegree:       MinDegree,
		maxElementCount: MinDegree * 2,
	}
	PBtree.maxSubCount = PBtree.maxElementCount + 1
	PBtree.minElementCount = MinDegree - 1
	PBtree.root = PBtree.nodeCreate(true)
	return nil
}

// Insert 用于向B树中插入一个元素
func (B *BTree) Insert(Key, Value string) {

}

// Delete 用于在B树中删除一个元素
func (B *BTree) Delete(Key string) {

}

// search 用于在B树中搜索指定的元素
func (B *BTree) search(Key string) string {
	PNewElement := &Element{key: Key}
	Pointer, Index := bTreeSearch(B.root, PNewElement)
	if Pointer == nil {
		return ""
	}
	return Pointer.eles[Index].val
}

// Min 返回B树中最小的元素
func (B *BTree) Min() {

}

// Max 返回B树中最大的元素
func (B *BTree) Max() {

}

// First 首元素
func (B *BTree) First() {

}

// Last 尾部元素
func (B *BTree) Last() {

}

// Next 用于下一个元素
func (B *BTree) Next() {

}

// Prev 用于上一个元素
func (B *BTree) Prev() {

}
