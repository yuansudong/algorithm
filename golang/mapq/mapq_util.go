package mapq

// Length 返回该MAPQ 的长度
func Length(MQ *MAPQ) int {
	return MQ.len
}

// Cap 返回该MAPQ 的容器量
func Cap(MQ *MAPQ) int {
	return MQ.cap
}

func parent(i int) int {
	return (i+1)>>1 - 1
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

// maxHeapify 用于建立最大堆性质
func maxHeapify(MQ *MAPQ, Index int) {
	L := left(Index)
	R := right(Index)
	MaxIndex := Index
	if L < MQ.len && MQ.store[L].compare(&MQ.store[MaxIndex]) == 1 {
		MaxIndex = L
	}
	if R < MQ.len && MQ.store[R].compare(&MQ.store[MaxIndex]) == 1 {
		MaxIndex = R
	}
	if MaxIndex != Index {
		MQ.store[Index], MQ.store[MaxIndex] = MQ.store[MaxIndex], MQ.store[Index]
		maxHeapify(MQ, MaxIndex)
	}

}

// 建立最大堆
func buildMaxHeap(MQ *MAPQ) {
	for i := MQ.len / 2; i >= 0; i-- {
		maxHeapify(MQ, i)
	}
}

// 堆排序
func heapSort(MQ *MAPQ) {
	buildMaxHeap(MQ)
	Len := MQ.len
	for i := MQ.len - 1; i > 0; i-- {
		MQ.store[0], MQ.store[i] = MQ.store[i], MQ.store[0]
		Len = Len - 1
		maxHeapifySort(MQ, 0, Len)
	}
}
func maxHeapifySort(MQ *MAPQ, Index, Length int) {
	L := left(Index)
	R := right(Index)
	MaxIndex := Index
	if L < Length && MQ.store[L].compare(&MQ.store[MaxIndex]) == 1 {
		MaxIndex = L
	}
	if R < Length && MQ.store[R].compare(&MQ.store[MaxIndex]) == 1 {
		MaxIndex = R
	}
	if MaxIndex != Index {
		MQ.store[Index], MQ.store[MaxIndex] = MQ.store[MaxIndex], MQ.store[Index]
		maxHeapifySort(MQ, MaxIndex, Length)
	}

}
