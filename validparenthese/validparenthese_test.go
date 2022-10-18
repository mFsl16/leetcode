package validparenthese

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidParenthese(t *testing.T) {

	testcase := []struct{
		Given string
		Expected bool
	} {
		{
			Given: "()",
			Expected: true,
		},{
			Given: "()[]{}",
			Expected: true,
		},{
			Given: "(]",
			Expected: false,
		},
		{
			Given: "]",
			Expected: false,
		},
	}

	for i := range testcase {
		tc := testcase[i]

		got := isValid(tc.Given)
		require.Equal(t, tc.Expected, got)
	}

}