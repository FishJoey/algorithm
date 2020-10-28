package binary_search

func SearchInsert(nums []int, target int) int {
	return search(nums, 0, len(nums)-1, target)
}

func search(nums []int, left, right, target int) int {
	if nums == nil {
		return 0
	}
	if target <= nums[left] {
		return left
	}
	if target > nums[right] {
		return right + 1
	}
	half := left + (right-left)/2
	if target <= nums[half] {
		return search(nums, left, half, target)
	} else {
		return search(nums, half+1, right, target)
	}
}


// 原始的插入时排序
func NaiveInsertSort(nums *[]int, target int) int {
	if len(*nums) == 0 {
		*nums = append(*nums, target)
		return 0
	}
	position := searchInsert(*nums, target)
	if position == len(*nums) {
		*nums = append(*nums, target)
		return position
	}

	*nums = append((*nums)[:position+1], (*nums)[position:]...)
	(*nums)[position] = target
	return position
}

func searchInsert(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, left, right, target int) int {
	if nums == nil {
		return 0
	}
	if target < nums[left] {
		return left
	}
	if target >= nums[right] {
		return right + 1
	}
	half := left + (right-left)/2
	if target <= nums[half] {
		return binarySearch(nums, left, half, target)
	} else {
		return binarySearch(nums, half+1, right, target)
	}
}
