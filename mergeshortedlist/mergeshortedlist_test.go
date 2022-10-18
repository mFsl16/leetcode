package mergeshortedlist

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMergeShortedList(t *testing.T) {

	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	result1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
		},
	}

	testcase := []struct{
		Given1 *ListNode
		Given2 *ListNode
		Expected *ListNode
	} {
		{
			Given1: list1,
			Given2: list2,
			Expected: result1,
		},
	}

	for i := range testcase {
		tc := testcase[i]
		got := mergeTwoLists(tc.Given1, tc.Given2)
		require.Equal(t, tc.Expected, got)
	}
}