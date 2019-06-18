package leetcode

import "algorithm/jz"

// Subtree of Another Tree
func isSubtree(root *jz.TreeNode, node *jz.TreeNode) bool {
	return jz.HasSubtree(root, node)
}
