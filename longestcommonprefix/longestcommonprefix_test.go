package longestcommonprefix

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestCommonPrefix(t *testing.T) {
	testcase := []struct {
		Given    []string
		Expected string
	}{
		{
			Given:    []string{"dog", "racecar", "car"},
			Expected: "",
		}, {
			Given:    []string{"flower", "flow", "flight"},
			Expected: "fl",
		}, {
			Given:    []string{"c", "c"},
			Expected: "c",
		},
	}

	for i := range testcase {
		tc := testcase[i]

		got := longestCommonPrefix(tc.Given)
		require.Equal(t, tc.Expected, got)
	}
}
