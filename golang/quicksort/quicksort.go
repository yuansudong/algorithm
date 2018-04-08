package quicksort

// QuickSort 用于快速排序
func QuickSort(Array []int, start, end int) {
	if start < end {
		Q := Partition(Array, start, end)
		QuickSort(Array, start, Q-1)
		QuickSort(Array, Q+1, end)
	}
}

// Partition 用于交换元素
func Partition(Array []int, start, end int) int {
	x := Array[end]
	i := start - 1
	for j := start; j < end; j++ {
		if Array[j] <= x {
			i = i + 1
			Array[i], Array[j] = Array[j], Array[i]
		}
	}
	Array[i+1], Array[end] = Array[end], Array[i+1]
	return i + 1
}
