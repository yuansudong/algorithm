package graph

const (
	// WHITE 白色
	WHITE = 0
	// GRAY 灰色
	GRAY = 1
	// BLACK 黑色
	BLACK = 2
	// WuQiong 表示无穷大
	WuQiong = 32768
)

type (
	// LVertix 用于表示一个顶点的信息
	LVertix struct {
		color    int
		distance int
		parent   *LVertix
		child    []*LVertix
	}
	// LinkedGraph 邻接链表
	LinkedGraph struct {
		// V用于表示这个图中的所有点
		V []*LVertix
		Q *Queue
	}
)

//BFS 广度优先搜索
func (LG *LinkedGraph) BFS(S *LVertix) {
	for _, TmpV := range LG.V {
		TmpV.color = WHITE
		TmpV.distance = WuQiong
		TmpV.parent = nil
	}
	S.color = GRAY
	S.distance = 0
	S.parent = nil
	LG.Q.EnQueue(S)
	for LG.Q.len != 0 {
		_, u := LG.Q.DeQueue()
		U := u.(*LVertix)
		for _, V := range U.child {
			if V.color == WHITE {
				V.color = GRAY
				V.distance = U.distance + 1
				V.parent = U
				LG.Q.EnQueue(V)
			}
			U.color = BLACK
		}

	}
}
