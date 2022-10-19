package removeelement

import "container/list"

func removeElement(nums []int, val int) int {

	stack := list.New().Init()

	for i := range nums {
		if nums[i] != val {
			stack.PushFront(nums[i])
		}
	}

	current := stack.Front()
	for j := stack.Len() - 1; j >= 0; j-- {
		nums[j] = current.Value.(int)
		current = current.Next()
	}

	return stack.Len()
}
