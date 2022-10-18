package removeduplicateshortedarray

import "container/list"

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	l := list.New().Init()

	for i := range nums {
		num := nums[i]

		if l.Front() == nil {
			l.PushFront(num)
			continue
		}

		if l.Front().Value != num {
			l.PushFront(num)
		}
	}

	val := l.Front()
	for i := l.Len() - 1; i >= 0; i-- {
		nums[i] = val.Value.(int)
		val = val.Next()
	}

	return l.Len()
}
