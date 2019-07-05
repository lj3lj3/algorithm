package jz

// 把二叉树打印成多行
//
// 从上到下按层打印二叉树，同一层结点从左至右输出。每一层输出一行。
//
// 使用计数器记录每个横向的node的数量 即当这一行遍历完之后 队列中存放的都是下一行的数据 而且下下一行还没有开始
func PrintIntoMultiLines(root *TreeNode) [][]int {
	listNodes := make([][]int, 0)
	queue := make([]*TreeNode, 0)

	// Push root to the queue
	queue = append(queue, root)

	listNodesRow := make([]int, 0)

	rowCount := 1

	for len(queue) > 0 {
		// Update row count
		if rowCount == 0 {
			// New line
			listNodesRow = make([]int, 0)
			rowCount = len(queue)
		}

		// Pop it
		node := queue[0]
		if len(queue) > 1 {
			queue = queue[1:] // New queue
		} else {
			queue = nil
		}

		if node == nil {
			continue
		}

		// Print value
		listNodesRow = append(listNodesRow, node.value)

		// Push left and right to the queue
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}

		rowCount--
		if rowCount == 0 {
			// Push to the list
			listNodes = append(listNodes, listNodesRow)
		}
	}

	return listNodes
}
