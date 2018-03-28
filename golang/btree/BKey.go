package btree

type (
	// BKey 是一个Key的结构
	BKey interface {
		compare(that *BKey) int
	}
)
