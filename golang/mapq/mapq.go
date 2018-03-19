package mapq

import (
	"fmt"
)

type (
	//MAPQ 为最大优先队列
	MAPQ struct {
		// len为当前已经有的长度
		len int
		// cap为容量
		cap int
		// store 为存储元素的地方
		store []element
	}
)

// Create 用于创建一个MAPQ
func Create() *MAPQ {
	PMAPQ := &MAPQ{}
	PMAPQ.store = make([]element, 1024)
	PMAPQ.len = 0
	PMAPQ.cap = 1024
	return PMAPQ
}

// Print 返回最大的元素
func (MQ *MAPQ) Print() {
	for i := 0; i < MQ.len; i++ {
		fmt.Println(MQ.store[i].value)
	}
	fmt.Println("***************")
}

// Max 返回最大的元素
func (MQ *MAPQ) Max() interface{} {
	return MQ.store[0].value
}

// ExtractMax 返回最大的元素并且删除该元素
func (MQ *MAPQ) ExtractMax() interface{} {
	if MQ.len != 0 {
		Max := MQ.store[0].value
		MQ.store[0] = MQ.store[MQ.len-1]
		MQ.len = MQ.len - 1
		maxHeapify(MQ, 0)
		return Max
	}
	return nil

}

// IncreaseKey 返回最大的元素并且删除该元素
func (MQ *MAPQ) IncreaseKey(Index int, Ele *element) {
	if Ele.compare(&MQ.store[Index]) == -1 {
		fmt.Println("111111")
		return
	}
	MQ.store[Index] = *Ele
	var p int
	for Index > 0 {
		p = parent(Index)
		if MQ.store[p].compare(&MQ.store[Index]) == -1 {
			MQ.store[p], MQ.store[Index] = MQ.store[Index], MQ.store[p]
			Index = p
			continue
		}
		break
	}
}

// Insert 返回最大的元素并且删除该元素
func (MQ *MAPQ) Insert(Order int, Value interface{}) {
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
func (MQ *MAPQ) Sort() {
	heapSort(MQ)
}
