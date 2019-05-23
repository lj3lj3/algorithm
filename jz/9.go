package jz

// 用两个栈实现队列
//
// 用两个栈来实现一个队列，完成队列的 Push 和 Pop 操作。
type TwoStackQueue struct {
	PushStack []int
	PopStack  []int
}

func (t *TwoStackQueue) Push(node int) {
	// transfer elements from pop stack to push tack first
	if len(t.PopStack) > 0 {
		for i := len(t.PopStack) - 1; i >= 0; i-- {
			t.PushStack = append(t.PushStack, t.PopStack[i])
		}
	}

	// reset pop stack
	t.PopStack = nil

	// push to stack
	t.PushStack = append(t.PushStack, node)
}

func (t *TwoStackQueue) Pop() int {
	var node int
	// transfer elements from push stack to pop stack first
	if len(t.PushStack) > 0 {
		for i := len(t.PushStack) - 1; i >= 0; i-- {
			t.PopStack = append(t.PopStack, t.PushStack[i])
		}
	}

	// pop last item
	length := len(t.PopStack)
	if length > 0 {
		node = t.PopStack[length-1]
		t.PopStack = t.PopStack[:length-1] // remove last item
	} else {
		// empty pop stack
		node = 0
	}

	// reset push stack
	t.PushStack = nil

	return node
}
