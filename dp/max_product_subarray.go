package dp

import (
	"algorithm/utils"
)

func MaxProductSubarray(nums []int) int {
	n := len(nums)
	dp0 := nums[0]
	dp1 := nums[0]
	var result = nums[0]
	for i := 1; i < n; i++ {
		v := nums[i]
		dp0, dp1 = utils.Max(dp0*v, dp1*v, v), utils.Min(dp1*v, dp0*v, v)
		result = utils.Max(dp0, dp1, result)
	}
	return result
}
