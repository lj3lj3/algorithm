package jz

import "math"

// 包含 min 函数的栈
//
// 定义栈的数据结构，请在该类型中实现一个能够得到栈最小元素的 min 函数。
//
// 使用双stack结构 一个stack是普通的stack 另一个是minStack 用于存储到当前位置为止的最小值
// 在push操作的使用生成当前位置的min值 然后push到minStack
type StackWithMinFunc struct {
	stack    []int
	minStack []int
}

func (s *StackWithMinFunc) push(node int) {
	// Push the node into the both stacks
	s.stack = append(s.stack, node)
	// Push the min value into the min stack
	value := 0
	if len(s.minStack) > 0 {
		value = int(math.Min(float64(s.min()), float64(node)))
	} else {
		value = node
	}
	s.minStack = append(s.minStack, value)
}

func (s *StackWithMinFunc) pop() int {
	if len(s.stack) == 0 {
		return 0
	}

	lastIndex := len(s.stack) - 1
	node := s.stack[lastIndex] // Get last node

	s.stack = s.stack[:lastIndex] // Remove last node
	s.minStack = s.minStack[:lastIndex]

	return node
}

func (s *StackWithMinFunc) top() int {
	if len(s.stack) == 0 {
		return 0
	}

	return s.stack[len(s.stack)-1]
}

func (s *StackWithMinFunc) min() int {
	if len(s.minStack) == 0 {
		return 0
	}

	return s.minStack[len(s.minStack)-1]
}
