package removeelement

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoveElement(t *testing.T) {

	testcase := []struct{
		GivenVal int
		GivenArr []int
		Expected int
	} {
		{
			GivenVal: 3,
			GivenArr: []int{3,2,2,3},
			Expected: 2,
		},
		{
			GivenVal: 2,
			GivenArr: []int{0,1,2,2,3,0,4,2},
			Expected: 5,
		},
	}

	for i := range testcase{
		tc := testcase[i]
		got := removeElement(tc.GivenArr, tc.GivenVal)

		require.Equal(t, tc.Expected, got)
	}
}
