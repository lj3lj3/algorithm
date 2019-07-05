package jz

// 按之字形顺序打印二叉树
//
// 请实现一个函数按照之字形打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右至左的顺序打印，第三行按照从左到右的顺序打印，其他行以此类推。
//
// 每行数据的打印上都需要进行相对应的反转操作
func PrintTreeInZhiFormat(root *TreeNode) [][]int {
	row := struct {
		nodesCount int
		index      int
	}{1, 0}

	listNodes := make([][]int, 0)
	queue := make([]*TreeNode, 0)

	// Push root to the queue
	queue = append(queue, root)

	listNodesRow := make([]int, 0)

	for len(queue) > 0 {
		// Update row count
		if row.nodesCount == 0 {
			// New line
			listNodesRow = make([]int, 0)
			row.index++
			row.nodesCount = len(queue)
			// Reverse the queue
			Reverse(queue)
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
		// EDIT: Push nodes in the queue based on index
		if row.index&1 == 0 {
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		} else {
			if node.right != nil {
				queue = append(queue, node.right)
			}
			if node.left != nil {
				queue = append(queue, node.left)
			}
		}

		row.nodesCount--
		if row.nodesCount == 0 {
			// Push to the list
			listNodes = append(listNodes, listNodesRow)
		}
	}

	return listNodes
}

func Reverse(nodes []*TreeNode) {
	for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	}
}
