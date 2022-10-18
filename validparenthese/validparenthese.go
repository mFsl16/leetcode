package validparenthese

import (
	"container/list"
	"strings"
)

func isValid(s string) bool {

	if len(s) == 0 {
		return false
	}

	brackets := strings.Split(s, "")
	stack := list.New()

	for i := range brackets {

		b := brackets[i]

		switch b {
		case "{", "(", "[" :
			stack.PushFront(b)
		case "}" :
			if stack.Front() != nil && stack.Front().Value == "{" {
				stack.Remove(stack.Front())
				break
			}
			stack.PushFront(b)
		case ")" :
			if stack.Front() != nil && stack.Front().Value == "(" {
				stack.Remove(stack.Front())
				break
			}
			stack.PushFront(b)
		case "]" :
			if stack.Front() != nil && stack.Front().Value == "[" {
				stack.Remove(stack.Front())
				break
			}
			stack.PushFront(b)
		}
	}

	return stack.Len() == 0
}

