package dp

// 找出和最大的子串
func MaxSubArray(nums []int) int {
	n := len(nums)
	result := nums[0]
	for i := 1; i < n; i++ {
		nums[i] = max(nums[i-1], 0) + nums[i]
		result = max(nums[i], result)
	}
	return result
}

func MaxSubArray2(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	result := nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1], 0) + nums[i]
		result = max(dp[i], result)
	}
	return result
}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
