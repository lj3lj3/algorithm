package jz

import "reflect"

// 二叉搜索树的后序遍历序列
//
// 输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。如果是则输出Yes,否则输出No。假设输入的数组的任意两个数字都互不相同。
//
// 后序遍历二叉树 然后和传入的数组进行对比
// 可以在每次append值的时候 对比当前传入的值是否与期望的值一致
func VerifySequenceOfBST(root *TreeNode, expected []int) bool {
	results := verify(root)
	return reflect.DeepEqual(results, expected)
}

func verify(node *TreeNode) []int {
	results := make([]int, 0)

	if node == nil {
		return results
	}

	results = append(results, verify(node.left)...)
	results = append(results, verify(node.right)...)

	return append(results, node.value)
}
