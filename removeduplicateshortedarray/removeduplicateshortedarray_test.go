package removeduplicateshortedarray

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicateShortedArray(t *testing.T) {

	testcase := []struct {
		Given    []int
		Expected int
	}{
		{
			Given:    []int{1, 1, 2},
			Expected: 2,
		},
		{
			Given:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			Expected: 5,
		},
	}

	for i := range testcase {
		tc := testcase[i]

		got := removeDuplicates(tc.Given)
		require.Equal(t, tc.Expected, got)
	}

}
