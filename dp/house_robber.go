package dp

import "algorithm/utils"

// 不能连续选择相邻的数字,找到最大的和
func houseRobber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = utils.Max(nums[0], nums[1])
	for i := 2; i < n; i += 1 {
		dp[i] = utils.Max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

// houses are arranged in circleR
// max(HouseRobber1([0...n-1]), HouseRobber1([1...n])
func HouseRobber2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return utils.Max(nums...)
	}

	n := len(nums)
	first := robberHelp(nums, 0, n-1)
	second := robberHelp(nums, 1, n)

	return utils.Max(first, second)
}

func robberHelp(nums []int, left, right int) int {
	dp := make([]int, right)
	dp[left] = nums[left]
	dp[left+1] = utils.Max(nums[left], nums[left+1])
	for i := left + 2; i < right; i++ {
		dp[i] = utils.Max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[right-1]
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// root是否被选
// dp[0] = dp.left + dp.right + n // root被选
// dp[1] = max(dp.left) + max(dp.right) // root 不被选
func HouseRobber3(root *TreeNode) int {
	result := dfs(root)
	return utils.Max(result[0], result[1])
}

func dfs(root *TreeNode) [2]int {
	if root == nil {
		return [2]int{0, 0}
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	currentRobbed := left[1] + root.Val + right[1]
	currentNotRobbed := utils.Max(left[0], left[1]) + utils.Max(right[0], right[1])
	return [2]int{currentRobbed, currentNotRobbed}
}
