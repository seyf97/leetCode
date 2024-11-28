package solutions

func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func RemoveElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	var indices []int = []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			indices = append(indices, i)
		}
	}

	for i := 0; i < len(nums); i++ {
		if i >= len(indices) {
			nums[i] = -1
		} else {
			nums[i] = nums[indices[i]]
		}
	}

	return len(indices)
}

func RemoveDuplicates(nums []int) int {
	uniqueVals := map[int]bool{}
	var new_nums []int = make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		_, exists := uniqueVals[nums[i]]
		if !exists {
			uniqueVals[nums[i]] = true
			new_nums = append(new_nums, nums[i])
		}
	}

	var dummyVals []int = make([]int, len(nums)-len(uniqueVals))

	for i, val := range append(new_nums, dummyVals...) {
		nums[i] = val
	}
	return len(uniqueVals)
}
