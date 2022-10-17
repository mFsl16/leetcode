package romantointeger

import (
	"container/list"
	"strings"
)

func romanToInt(s string) int {

	romans := strings.Split(s, "")
	stack := list.New()

	sum := 0

	for i := range romans {
		roman := romans[i]
		switch roman {
		case "I":
			sum += 1
			stack.PushFront("I")
		case "V":
			if stack.Front() != nil && stack.Front().Value == "I" {
				sum -= 1
				sum += 4
				break
			} else {
				sum += 5
			}
			stack.PushFront("V")

		case "X":
			if stack.Front() != nil && stack.Front().Value == "I" {
				sum -= 1
				sum += 9
			} else {
				sum += 10
			}
			stack.PushFront("X")

		case "L":
			if stack.Front() != nil && stack.Front().Value == "X" {
				sum -= 10
				sum += 40
			} else {
				sum += 50
			}
			stack.PushFront("L")
		case "C":
			if stack.Front() != nil && stack.Front().Value == "X" {
				sum -= 10
				sum += 90
			} else {
				sum += 100
			}
			stack.PushFront("C")
		case "D":
			if stack.Front() != nil && stack.Front().Value == "C" {
				sum -= 100
				sum += 400
			} else {
				sum += 500
			}
			stack.PushFront("D")
		case "M":
			if stack.Front() != nil && stack.Front().Value == "C" {
				sum -= 100
				sum += 900
			} else {
				sum += 1000
			}
			stack.PushFront("M")
		}
	}

	return sum
}
