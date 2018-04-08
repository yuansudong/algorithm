package btree

func bTreeSearch(Node *bnode, Ele *Element) (*bnode, int) {
	CompareRet, i := 0, 0
	for ; i < Node.size; i++ {
		CompareRet = Ele.Compare(Node.eles[i])
		if CompareRet != 1 {
			break
		}
	}
	if CompareRet == 0 {
		return Node, i
	}
	if Node.isLeaf {
		return nil, i
	}
	return bTreeSearch(Node.sub[i], Ele)
}

// bTreeSplitChild 用于分裂一个新的兄弟结点
func bTreeSplitChild(B *BTree, Node *bnode, i int) {
	y := Node.sub[i]
	z := B.nodeCreate(y.isLeaf)
	z.size = B.minDegree - 1
	for j := 0; j < z.size; j++ {
		z.eles[i] = y.eles[j+z.size]
	}
	y.size = B.minDegree - 1
	for j := Node.size + 1; j >= i+1; j-- {
		Node.sub[j+1] = Node.sub[j]
	}
	Node.sub[i+1] = z
	for j := Node.size - 1; j >= i; j-- {
		Node.eles[j+1] = Node.eles[j]
	}
	Node.eles[i] = y.eles[B.minDegree]
	Node.size++
}

// bTreeGetPosition
func bTreeGetPosition(PNode *bnode, Ele *Element) (*bnode, int) {
	i, compareRet := 0, 0
	for ; i < PNode.size; i++ {
		compareRet = Ele.Compare(PNode.eles[i])
		if compareRet != 1 {
			break
		}
	}
	if PNode.isLeaf {
		return PNode, i
	}
	return bTreeGetPosition(PNode.sub[i], Ele)
}

// bTree
func bTreeInsertChildNode(PNode *bnode, PInsertEle *Element) int {
	i, CompareRet := 0, 0
	for ; i < PNode.size; i++ {
		CompareRet = PInsertEle.Compare(PNode.eles[i])
		if CompareRet != 1 {
			break
		}
	}
	return i
}

// bTreeSplit 用于分裂子树
func bTreeSplit(B *BTree, SplitNode *bnode) {

	// 如果满足该条件则代表着是要进行分裂了
	if SplitNode.size == B.maxElementCount {
		// 代表分裂的是叶子结点
		i, length := 0, 0
		NewSplitNode := B.nodeCreate(SplitNode.isLeaf)
		/* 构建新分裂结点 */
		// 处理新元素结点
		i = B.minDegree + 1
		length = SplitNode.size
		for ; i < length; i++ {
			NewSplitNode.eles[NewSplitNode.size] = SplitNode.eles[i]
			NewSplitNode.size++
		}
		// 处理新元素的子结点
		i = B.minDegree + 1
		length = SplitNode.size + 1
		for j := 0; i < length; i++ {
			NewSplitNode.sub[j] = SplitNode.sub[i]
			j++
		}
		/* 重新构建分裂结点 */
		// 构建分裂结点的
		i = SplitNode.size - 1
		for ; i >= B.minDegree; i-- {
			SplitNode.eles[i] = nil
			SplitNode.size--
		}
		// 构建分裂结点的子结点
		i = B.minDegree + 1
		for ; i < B.maxElementCount; i++ {
			SplitNode.sub[i] = nil
		}
		// 构建父结点
		ParentNode := SplitNode.parent
		if ParentNode == nil {
			// 如果进入这里,就代表父结点是空的
			ParentNode := B.nodeCreate(false)
			B.root = ParentNode
		}
		ParentInsertElement := SplitNode.eles[B.minDegree]
		Index := bTreeInsertChildNode(ParentNode, ParentInsertElement)
		ParentNode.sub[Index] = SplitNode
		ParentNode.sub[Index+1] = NewSplitNode
		SplitNode.parent = ParentNode
		NewSplitNode.parent = ParentNode
		bTreeSplit(B, ParentNode)
	}

}

// bTreeInsert 用于向B树中插入一个元素
func bTreeInsert(B *BTree, Ele *Element) {
	// 找到位置
	InsertNode, Index := bTreeGetPosition(B.root, Ele)
	// 在该位置保存元素的指针
	for j := InsertNode.size - 1; j > Index; j-- {
		InsertNode.eles[j+1] = InsertNode.eles[j]
	}
	InsertNode.eles[Index] = Ele
	InsertNode.size++
	bTreeSplit(B, InsertNode)
}

// bTreeMergeNode(MinChildNode, MaxChildNode)
func bTreeMergeNode(B *BTree, MinChildNode, MaxChildNode *bnode) bool {
	if MinChildNode.size+MaxChildNode.size >= B.maxElementCount {

	}
}

// bTreeDelete 用于从B树中删除一个元素
func bTreeDelete(B *BTree, Ele *Element) {
	PDeleteNode, Index := bTreeSearch(B.root, Ele)
	if PDeleteNode != nil {
		// 如果是叶结点
		i := 0
		length := 0
		if PDeleteNode.isLeaf {
			i = Index
			length = PDeleteNode.size - 1
			for ; i < length; i++ {
				PDeleteNode.eles[i] = PDeleteNode.eles[i+1]
			}
			PDeleteNode.eles[length] = nil
			PDeleteNode.size = length
		} else {
			MinChildNode := PDeleteNode.sub[Index]
			MaxChildNode := PDeleteNode.sub[Index+1]
			MergeChildNode := bTreeMergeNode(MinChildNode, MaxChildNode)

		}

	}
}

// bTreeDeleteCheck 用于删除元素之后进行检查
func bTreeDeleteCheck(B *BTree, PDeleteNode *bnode, MinSubIndex int) {

	if !PDeleteNode.isLeaf {
		// 如果不是叶子结点则需要进行处理
		MinSubNode := PDeleteNode.sub[MinSubIndex]
		MaxSubNode := PDeleteNode.sub[MinSubIndex+1]

		if MinSubNode.size+MaxSubNode.size >= B.maxElementCount {
			// 如果条件满足那么就需要在最多元素中删除一个元素,将其替换到原来的位置

		}

	}

}
