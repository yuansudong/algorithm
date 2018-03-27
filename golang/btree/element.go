package btree

// Element 用于表示一个元素
type Element struct {
	key   interface{}
	value interface{}
}

// CreateElement 用于创建元素
func CreateElement(Key, Value interface{}) *Element {
	return &Element{
		key:   Key,
		value: Value,
	}
}
