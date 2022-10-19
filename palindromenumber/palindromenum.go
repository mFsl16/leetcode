package palindromenumber

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	// 123

	reversedX := 0
	notReversed := x

	for x != 0 {
		num := x % 10
		reversedX = reversedX*10 + num
		x = x / 10
	}

	return reversedX == notReversed
}
