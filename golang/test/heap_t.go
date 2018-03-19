package main

type (
	// Heap 用于描述一个堆
	Heap struct {
		// 堆的所存储的元素个数
		len int
		// 堆的容量
		cap int
		//
		array []int
	}
)

// 获得父结点的下标
func parent(i int) int {
	return ((i+1)<<1)/2 - 1
}

// 0   , 1,2,3,4,5,6
// 获得左结点的下标
func left(i int) int {
	return (i+1)<<1 - 1
}

// 获得右结点的下标
func right(i int) int {
	return (i + 1) << 1
}

// Create 用于创建一个堆
func Create(A []int) *Heap {
	PHeap := &Heap{}
	PHeap.len = len(A)
	PHeap.array = A
	return PHeap
}

// compare 用于比较堆内的元素,-1 为小于,0为等于,1为大于
func compare(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	}
	return -1
}

// maxHeapfiy 用于维护最大堆的性质
func maxHeapfiy(H *Heap, I int) {
	// 该函数用于维护最大堆的性质
	MaxIndex := I
	L := left(I)
	R := right(I)
	if L < H.len && H.array[L] > H.array[I] {
		MaxIndex = L
	} else {
		MaxIndex = I
	}
	if R < H.len && H.array[R] > H.array[MaxIndex] {
		MaxIndex = R
	}
	if MaxIndex != I {
		Temp := H.array[I]
		H.array[I] = H.array[MaxIndex]
		H.array[MaxIndex] = Temp
		// 因为调换了相关位置元素的值, 这样可能引起,MaxIndex为根结点的不符合最大堆的性质
		// 所以需要递归调用
		maxHeapfiy(H, MaxIndex)
	}
}

// createMaxHeap 用于创建一个最大堆
func createMaxHeap(H *Heap) {
	H.len = len(H.array)
	for I := H.len >> 1; I >= 0; I-- {
		maxHeapfiy(H, I)
	}
}

// heapSort 用于对一个堆进行排序 nlgn
func heapSort(H *Heap) {
	// 首先建立最大堆
	createMaxHeap(H)

	var Temp int
	for I := H.len - 1; I >= 1; I-- {
		Temp = H.array[I]
		H.array[I] = H.array[0]
		H.array[0] = Temp
		H.len--
		maxHeapfiy(H, 0)
	}

}

// GetSet 用于获取有序的集合
func GetSet(H *Heap) []int {
	return H.array
}
