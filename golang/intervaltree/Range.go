package intervaltree

type (
	// Range 表示的是一个区间
	Range struct {
		high int
		low  int
	}
)

// CreateRange 用于创建一个区间
func CreateRange(low, high int) *Range {
	return &Range{
		low:  low,
		high: high,
	}
}

// GetHigh 获取高界限
func (r *Range) GetHigh() int {
	return r.high
}

// GetLow 获取低界线
func (r *Range) GetLow() int {
	return r.low
}

// Overlap 用于判断两个区间是否重叠
func (r *Range) Overlap(t *Range) bool {
	if r.low <= t.high && r.high >= t.low {
		return true
	}
	return false
}

// min 用于求出最小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max 用于求出最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
