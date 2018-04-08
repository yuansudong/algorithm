package btree

type (
	// Element 用于表示BTree 中的元素
	Element struct {
		key string
		val string
	}
)

// CreateElement 用于创建一个BTree中的元素
func (E *Element) CreateElement(Key, Val string) *Element {
	return &element{
		key: Key,
		val: Val,
	}
}

// Compare 用于比较两个BTree元素 0等于,-1小于,1大于
func (E *Element) Compare(T *Element) int {
	if E.key < T.key {
		return -1
	} else if E.key == T.key {
		return 0
	}
	return 1
}
