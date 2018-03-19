package slice

type (
	// Slice 用于描述一个切片的结构
	Slice struct {
		// len为当前已经有的长度
		len int
		// cap为容量
		cap int
		// store 为存储元素的地方
		store []interface{}
	}
)

// Create 用于创建一个slice
func Create() *Slice {
	PSlice := &Slice{}
	PSlice.store = make([]interface{}, 1024)
	PSlice.len = 0
	PSlice.cap = 1024
	return PSlice
}

// Get 用于获取指定索引的值
func (S *Slice) Get(i int) interface{} {
	if i < S.len {
		return S.store[i]
	}
	return nil
}

// Set 用于设置一个值
func (S *Slice) Set(i int, Value interface{}) {
	if i < S.len {
		S.store[i] = Value
		return
	}
}

// Length 返回该Slice 的长度
func (S *Slice) Length() int {
	return S.len
}

// Cap 返回该Slice 的容器量
func (S *Slice) Cap() int {
	return S.cap
}

// Append 用于向Slice 末尾追加函数
func (S *Slice) Append(Value interface{}) {
	if S.len < S.cap {
		S.store[S.len] = Value
		S.len++
		return
	}
	// 如果执行到了这里,代表内存不够
	S.cap += S.cap / 4
	OldSlice := S.store
	S.store = make([]interface{}, S.cap)
	for i := 0; i < S.len; i++ {
		S.store[i] = OldSlice[i]
	}
	S.Append(Value)
}

// Delete 用于删除并且返回该元素
func (S *Slice) Delete(i int) (Ret interface{}) {
	if i < S.len {
		Ret = S.store[i]
		for J := i + 1; J < S.len; J++ {
			S.store[J-1] = S.store[J]
		}
		S.len = S.len - 1
		// 因为删除了元素,所以要判断当前的容量
		MaxCap := S.len + S.len/4
		if MaxCap < S.cap && S.len > 1024 {
			// 如果条件满足那么开始缩减容器的大小
			S.cap = MaxCap
			OldSlice := S.store
			S.store = make([]interface{}, S.cap)
			for i := 0; i < S.len; i++ {
				S.store[i] = OldSlice[i]
			}
		}
	}
	return Ret

}

// Swap 用于交换两个索引的值
func (S *Slice) Swap(this, that int) {
	if this < S.len && that < S.len {
		Temp := S.store[this]
		S.store[this] = S.store[that]
		S.store[that] = Temp
	}

}

/*---------inner method-----------*/
