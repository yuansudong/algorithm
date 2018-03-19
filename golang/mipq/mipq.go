package mipq

import (
	"fmt"
)

type (
	//MIPQ 为最大优先队列
	MIPQ struct {
		// len为当前已经有的长度
		len int
		// cap为容量
		cap int
		// store 为存储元素的地方
		store []element
	}
)

// Create 用于创建一个MIPQ
func Create() *MIPQ {
	PMAPQ := &MIPQ{}
	PMAPQ.store = make([]element, 1024)
	PMAPQ.len = 0
	PMAPQ.cap = 1024
	return PMAPQ
}

// Print 返回最大的元素
func (MQ *MIPQ) Print() {
	for i := 0; i < MQ.len; i++ {
		fmt.Println(MQ.store[i].value)
	}
	fmt.Println("***************")
}

// Min 返回最小的元素
func (MQ *MIPQ) Min() interface{} {
	return MQ.store[0].value
}

// ExtractMin 返回最小的元素并且删除该元素
func (MQ *MIPQ) ExtractMin() interface{} {
	if MQ.len != 0 {
		Min := MQ.store[0].value
		MQ.store[0] = MQ.store[MQ.len-1]
		MQ.len = MQ.len - 1
		minHeapify(MQ, 0)
		return Min
	}
	return nil

}

// IncreaseKey 返回最大的元素并且删除该元素
func (MQ *MIPQ) IncreaseKey(Index int, Ele *element) {
	if Ele.compare(&MQ.store[Index]) == -1 {
		fmt.Println("111111")
		return
	}
	MQ.store[Index] = *Ele
	var p int
	for Index > 0 {
		p = parent(Index)
		if MQ.store[p].compare(&MQ.store[Index]) == 1 {
			MQ.store[p], MQ.store[Index] = MQ.store[Index], MQ.store[p]
			Index = p
			continue
		}
		break
	}
}

// Insert 返回最大的元素并且删除该元素
func (MQ *MIPQ) Insert(Order int, Value interface{}) {
	if MQ.len < MQ.cap {
		NewEle := element{
			order: Order,
			value: Value,
		}
		MQ.store[MQ.len].order = 0
		MQ.IncreaseKey(MQ.len, &NewEle)
		MQ.len = MQ.len + 1
	}
}

// Sort 对其进行排序
func (MQ *MIPQ) Sort() {
	heapSort(MQ)
}
