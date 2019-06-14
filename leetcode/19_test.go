package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	// construct link list
	tests := []struct {
		test *ListNode
		k    int
		want *ListNode
	}{
		{
			test: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
				},
			},
			k: 2,
			want: &ListNode{
				Val: 1,
			},
		},
		{
			test: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 3,
							},
						},
					},
				},
			},
			k: 4,
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
		},
		{
			test: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 4,
								},
							},
						},
					},
				},
			},
			k: 2,
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
							},
						},
					},
				},
			},
		},
	}

	for _, value := range tests {
		if result := RemoveNthFromEnd(value.test, value.k); !reflect.DeepEqual(result, value.want) {
			t.Errorf("RemoveNthFromEnd failed, test: %v, k: %d, result: %v, want: %v", value.test, value.k, result, value.want)
		}
	}
}
