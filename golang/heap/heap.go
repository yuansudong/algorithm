package heap

// compare 用于比较堆内的元素,-1 为小于,0为等于,1为大于
func compare(a, b interface{}) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	}
	return -1
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
