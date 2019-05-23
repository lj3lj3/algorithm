package jz

import "testing"

const actionPush int8 = 1
const actionPop int8 = -1

type action struct {
	action int8
	data   int
}

func TestTwoStackQueue(t *testing.T) {
	tests := []struct {
		test []*action
		want []int
	}{
		{
			test: []*action{
				{
					action: actionPush,
					data:   2,
				},
				{
					action: actionPush,
					data:   4,
				},
				{
					action: actionPush,
					data:   80,
				},
				{
					action: actionPop,
				},
				{
					action: actionPop,
				},
				{
					action: actionPop,
				},
				{
					action: actionPop,
				},
				{
					action: actionPush,
					data:   20,
				},
				{
					action: actionPush,
					data:   26,
				},
				{
					action: actionPop,
				},
			},
			want: []int{2, 4, 80, 0, 20},
		},
		{
			test: []*action{
				{
					action: actionPush,
					data:   -1,
				},
				{
					action: actionPush,
					data:   0,
				},
				{
					action: actionPush,
					data:   -2,
				},
				{
					action: actionPop,
				},
				{
					action: actionPop,
				},
				{
					action: actionPop,
				},
				{
					action: actionPop,
				},
				{
					action: actionPush,
					data:   2,
				},
				{
					action: actionPush,
					data:   2,
				},
				{
					action: actionPop,
				},
				{
					action: actionPop,
				},
			},
			want: []int{-1, 0, -2, 0, 2, 2},
		},
		{
			test: []*action{
				{
					action: actionPop,
				},
			},
			want: []int{0},
		},
	}

	for _, value := range tests {
		queue := new(TwoStackQueue)
		popIndex := 0
		for _, test := range value.test {
			if test.action == actionPush {
				queue.Push(test.data)
			} else {
				if result := queue.Pop(); result != value.want[popIndex] {
					// failed
					t.Errorf("Pop failed, test: %v, result: %d, want: %d", value.test, result, value.want[popIndex])
				}
				popIndex++
			}
		}
	}
}
