package mipq

type element struct {
	// 用于表示顺序
	order int
	value interface{}
}

// compare 用于比较元素的值,-1 代表小于,0代表等于,1代表大于
func (E *element) compare(that *element) int {
	if E.order < that.order {
		return -1
	} else if E.order == that.order {
		return 0
	}
	return 1
}
