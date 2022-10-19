package searchinsertposition

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchInsertPosition(t *testing.T) {

	testcase := []struct {
		GivenArr    []int
		GivenTarget int
		Expected    int
	}{
		{
			GivenArr:    []int{1, 3, 5, 6},
			GivenTarget: 5,
			Expected:    2,
		},
		{
			GivenArr:    []int{1, 3, 5, 6},
			GivenTarget: 2,
			Expected:    1,
		},
		{
			GivenArr:    []int{1, 3, 5, 6},
			GivenTarget: 7,
			Expected:    4,
		}, {
			GivenArr:    []int{1, 3, 5, 6},
			GivenTarget: 0,
			Expected:    0,
		},
		{
			GivenArr:    []int{1},
			GivenTarget: 0,
			Expected:    0,
		},
	}

	for i := range testcase {
		tc := testcase[i]

		got := searchInsert(tc.GivenArr, tc.GivenTarget)
		require.Equal(t, tc.Expected, got)
	}
}
