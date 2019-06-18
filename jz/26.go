package jz

// 树的子结构
//
// 输入两棵二叉树A，B，判断B是不是A的子结构。(我们约定空树不是任意一个树的子结构)
func HasSubtree(root, node *TreeNode) bool {
	// Check if root or node is nil
	if root == nil || node == nil {
		return false
	}

	// Check if current node is starter of the subtree, otherwise move to the leaves
	return isSubTreeOfRoot(root, node) || HasSubtree(root.left, node) || HasSubtree(root.right, node)
}

func isSubTreeOfRoot(root, node *TreeNode) bool {
	if node == nil && root == nil {
		// node is nil and root is nil
		return true
	} else if root == nil || node == nil {
		// root is nil or node is nil
		// The leaf must be same
		return false
	}

	// Check value
	if root.value != node.value {
		// These nodes are different
		return false
	}

	// These nodes are equals to each other, check left leaf and right leaf
	return isSubTreeOfRoot(root.left, node.left) && isSubTreeOfRoot(root.right, node.right)
}
