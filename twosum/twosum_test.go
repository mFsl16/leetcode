package twosum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSum(t *testing.T) {

	testcase := []struct {
		Target   int
		Nums     []int
		ExResult []int
	}{
		{
			Target:   9,
			Nums:     []int{2, 7, 11, 15},
			ExResult: []int{0, 1},
		},
		{
			Target:   6,
			Nums:     []int{3, 2, 4},
			ExResult: []int{1, 2},
		},
		{
			Target:   6,
			Nums:     []int{3, 3},
			ExResult: []int{0, 1},
		},
		{
			Target:   4,
			Nums:     []int{2, 3, 5},
			ExResult: nil,
		},
	}

	for i := range testcase {
		tc := testcase[i]

		require.Equal(t, tc.ExResult, Twosum(tc.Nums, int(tc.Target)))
	}

}
