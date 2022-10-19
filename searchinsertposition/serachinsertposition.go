package searchinsertposition

func searchInsert(nums []int, target int) int {

	i := 0

	for i < len(nums) {
		currentNum := nums[i]
		if currentNum == target {
			return i
		} else if i == len(nums)-1 && currentNum < target {
			return i + 1
		} else if currentNum < target && nums[i+1] > target {
			return i + 1
		} else if currentNum > target {
			return i
		}
		i++
	}

	return 0
}
