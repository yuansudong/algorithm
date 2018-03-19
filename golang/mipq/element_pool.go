package mipq

const (
	defaultElementPoolSize = 256
)

type elementPool struct {
	start int
	end   int
	pool  []element
}

// createElementPool 用于创建一个元素池
func createElementPool() *elementPool {
	PElementPool := &elementPool{
		start: 0,
		end:   defaultElementPoolSize,
		pool:  make([]element, defaultElementPoolSize),
	}
	return PElementPool
}

// getElement 用于获取一个元素
func (E *elementPool) getElement() *element {
	if E.start < E.end {
		NewElement := &E.pool[E.start]
		E.start = E.start + 1
		return NewElement
	}
	E.pool = make([]element, defaultElementPoolSize)
	E.start = 0
	E.end = defaultElementPoolSize
	return E.getElement()
}
