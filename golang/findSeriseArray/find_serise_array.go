package findSeriseArray

// singleArrayFind 单个数组查找
func singleArrayFind(Array []int, low int, mid int, high int) (int, int, int) {
	// 左数组中最大的下标
	var leftMax int
	// 左数组中的最大和
	var leftSum int = -32768

	// 右数组中最大的下标
	var rightMax int
	// 右数组最大的和
	var rightSum int = -32768

	// sum为一个中间求和内存
	var sum int
	for i := mid; i >= 0; i-- {
		sum += Array[i]
		if sum > leftSum {
			leftSum = sum
			leftMax = i
		}
	}
	sum = 0
	for j := mid + 1; j <= high; j++ {
		sum += Array[j]
		if sum > rightSum {
			rightSum = sum
			rightMax = j
		}
	}
	return leftMax, rightMax, leftSum + rightSum
}

// FindSumMaxArray 查找和最大的子数组
func FindSumMaxArray(Array []int, low int, high int) (int, int, int) {
	if low == high {
		return low, high, Array[low]
	}

	mid := (low + high) / 2
	leftLow, leftHigh, leftSum := FindSumMaxArray(Array, low, mid)
	rightLow, rightHigh, rightSum := FindSumMaxArray(Array, mid+1, high)
	currentLow, currentHigh, currentSum := singleArrayFind(Array, low, mid, high)

	if leftSum >= rightSum && leftSum >= currentSum {
		return leftLow, leftHigh, leftSum
	} else if rightSum >= leftSum && rightSum >= currentSum {
		return rightLow, rightHigh, rightSum
	} else {
		return currentLow, currentHigh, currentSum
	}

}
