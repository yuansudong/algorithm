package vebtree

type (
	// VebNode 用于描述一个Veb的结点
	VebNode struct {
		u, min, max int
		summary     []byte
		cluster     []*VebNode
	}
)

// Min 找出最小值
func Min(V *VebNode) int {
	return V.min
}

// Max 找出最大值
func Max(V *VebNode) int {
	return V.max
}

// IsMember 用于判断x,是为为VebNode的元素
