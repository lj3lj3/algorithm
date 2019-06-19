package jz

// 对称的二叉树
//
// 请实现一个函数，用来判断一颗二叉树是不是对称的。注意，如果一个二叉树同此二叉树的镜像是同样的，定义其为对称的。
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// Check left and right leaf is same or not
	return isSymmetricSub(root.left, root.right)
}

func isSymmetricSub(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	}

	// Check value
	return left.value == right.value && isSymmetricSub(left.left, right.right) && isSymmetricSub(left.right, right.left)
}
