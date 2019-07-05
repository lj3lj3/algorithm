package jz

// 从上往下打印二叉树
//
// 从上往下打印出二叉树的每个节点，同层节点从左至右打印。
//
// 本质是广度优先遍历 使用队列依次存储每个节点
func PrintFromTopToBottom(root *TreeNode) []int {
	listNodes := make([]int, 0)
	queue := make([]*TreeNode, 0)

	// Push root to the queue
	queue = append(queue, root)

	for len(queue) > 0 {
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
		listNodes = append(listNodes, node.value)

		// Push left and right to the queue
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}

	return listNodes
}
