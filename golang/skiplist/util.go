package skiplist

import "math/rand"

// randomLevel 用于随机出一个层数
func randomLevel() (retLevel int) {
	level := 1
	for {
		if rand.Intn(0xFFFF) < skipRandomCrisis {
			goto end
		}
		level++
	}
end:
	if level < skipListMaxLevel {
		retLevel = level
	} else {
		retLevel = skipListMaxLevel - 1
	}
	return retLevel
}
