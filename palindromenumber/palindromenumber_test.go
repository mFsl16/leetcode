package palindromenumber

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPalindromeNum(t *testing.T) {

	testcase := []struct {
		Given  int
		Expect bool
	}{
		{
			Given:  121,
			Expect: true,
		},
		{
			Given:  -121,
			Expect: false,
		},
		{
			Given:  10,
			Expect: false,
		},
	}

	for i := range testcase {
		tc := testcase[i]

		got := isPalindrome(tc.Given)

		require.Equal(t, tc.Expect, got)
	}
}
