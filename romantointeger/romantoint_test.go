package romantointeger

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRomanToInt(t *testing.T) {

	testcase := []struct {
		Given    string
		Expected int
	}{
		{
			Given:    "III",
			Expected: 3,
		},
		{
			Given:    "LVIII",
			Expected: 58,
		},
		{
			Given:    "MCMXCIV",
			Expected: 1994,
		},
	}

	for i := range testcase {

		tc := testcase[i]

		got := romanToInt(tc.Given)
		require.Equal(t, tc.Expected, got)
	}
}
