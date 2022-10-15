package twosum

func Twosum(nums []int, target int) []int {

	store := make(map[int]int)

	for i := range nums {

		diff := target - nums[i]
		if _, ok := store[diff]; ok {
			return []int{store[diff], i}
		}
		store[nums[i]] = i
	}

	return nil
}
