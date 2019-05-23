package leetcode

import "algorithm/jz"

// 232. Implement Queue using Stacks
type TwoStackQueue struct {
	jz.TwoStackQueue
}

func (t *TwoStackQueue) Peek() int {
	var node int
	// transfer elements from push stack to pop stack first
	if len(t.PushStack) > 0 {
		for i := len(t.PushStack) - 1; i >= 0; i-- {
			t.PopStack = append(t.PopStack, t.PushStack[i])
		}
	}

	// reset push stack
	t.PushStack = nil

	// peek
	length := len(t.PopStack)
	if length > 0 {
		node = t.PopStack[length-1]
	} else {
		// empty pop stack
		node = 0
	}

	return node
}

func (t *TwoStackQueue) Empty() bool {
	return len(t.PopStack) == 0 && len(t.PushStack) == 0
}
