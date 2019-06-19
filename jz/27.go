package jz

// 操作给定的二叉树，将其变换为源二叉树的镜像。
//
// 二叉树的镜像定义：源二叉树
//      8
//     /  \
//    6   10
//   / \  / \
//  5  7 9  11
//
//  镜像二叉树
//      8
//     /  \
//    10   6
//   / \  / \
//  11 9 7   5
func Mirror(node *TreeNode) {
	if node == nil {
		return
	}

	// swap left leaf and right leaf
	node.left, node.right = node.right, node.left
	Mirror(node.left)
	Mirror(node.right)
}
