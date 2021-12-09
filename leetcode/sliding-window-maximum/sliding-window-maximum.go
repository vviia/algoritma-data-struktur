package slidingwindowmaximum

func maxSlidingWindow(nums []int, k int) []int {
	size := len(nums)
	if k <= 1 {
		return nums
	}

	g := k - 1

	left := make([]int, size)
	for i := 0; i < size; i++ {
		if i%g == 0 {
			left[i] = nums[i]
		} else {
			left[i] = max(nums[i], left[i-1])
		}
	}

	right := make([]int, size)
	right[size-1] = nums[size-1]
	for j := size - 2; j >= 0; j-- {
		if (j+1)%g == 0 {
			right[j] = nums[j]
		} else {
			right[j] = max(nums[j], right[j+1])
		}
	}

	res := make([]int, size-k+1)
	for i := 0; i <= size-k; i++ {
		res[i] = max(right[i], left[i+k-1])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
