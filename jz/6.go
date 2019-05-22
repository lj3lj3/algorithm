package jz

type list struct {
	value string
	next  *list
}

// 从尾到头打印链表
//
// 从尾到头反过来打印出每个结点的值。
func printListFromTailToHead(list *list) string {
	// using stack
	var stack []string
	for {
		stack = append(stack, list.value)
		if list.next == nil {
			// last one
			break
		} else {
			list = list.next
		}
	}

	// print
	var str string
	for i := len(stack) - 1; i >= 0; i-- {
		str += stack[i] + ","
	}
	return str[:len(str)-1]
}
