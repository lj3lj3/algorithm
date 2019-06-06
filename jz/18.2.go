package jz

type LinkNode struct {
	value int
	next  *LinkNode
}

// 18.2 删除链表中重复的结点
// 在一个排序的链表中，存在重复的结点，请删除该链表中重复的结点，重复的结点不保留，返回链表头指针。 例如，链表1->2->3->3->4->4->5 处理后为 1->2->5
//
// 删除链表中的node
//   node.next = node.next.next
// 移动当前指向到下一个node
//   pointer = pointer.next
func DeleteDuplication(nodeHead *LinkNode) *LinkNode {
	// check if link node is valid
	if nodeHead == nil || nodeHead.next == nil {
		return nodeHead
	}

	// set head node when there is first not equals exists
	firstNotEquals := true
	// init empty starter node
	// this will treat all nodes as this node's next nodes
	pointer := &LinkNode{}
	pointer.next = nodeHead

	for pointer.next != nil && pointer.next.next != nil {
		// check if next node is same as next's next node
		if pointer.next.value == pointer.next.next.value {
			// equals, remove next nodes which value is equals to this node's value
			value := pointer.next.value
			for pointer.next != nil && pointer.next.value == value {
				pointer.next = pointer.next.next
			}
		} else {
			if firstNotEquals {
				nodeHead = pointer.next // make nodeHead point to this node
				firstNotEquals = false
			}
			// not equals, move next
			pointer = pointer.next
		}
	}

	// check if there is is all equals in the list
	if firstNotEquals {
		nodeHead = nil
	}

	return nodeHead
}
