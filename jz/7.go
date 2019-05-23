package jz

import "strconv"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func (t *TreeNode) String() string {
	str := strconv.Itoa(t.value) + "|"
	if t.left != nil {
		str += t.left.String()
	}
	if t.right != nil {
		str += t.right.String()
	}
	return str
}

// 重建二叉树
//
// 根据二叉树的前序遍历和中序遍历的结果，重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
//
// 前序和中序遍历指的是root的位置
//
// 前序遍历的值在中序遍历中将中序遍历分割成两部分 中序遍历左边的为左子树 右面的为右子树 **前序遍历+中序遍历左子树的数量=前序遍历右子树的位置**
func ReConstructBinaryTree(preOrder, midOrder []int) *TreeNode {
	if len(preOrder) == 0 || len(midOrder) == 0 {
		return nil
	}

	parentNode := new(TreeNode)
	constructTreeNodes(preOrder, 0, parentNode, midOrder)

	return parentNode
}

func constructTreeNodes(preOrder []int, index int, parentNode *TreeNode, midOrder []int) {
	if index == len(preOrder) {
		// finished
		return
	}
	// split
	for i, midOrderValue := range midOrder {
		if midOrderValue == preOrder[index] {
			// found
			parentNode.value = midOrderValue
			// split it
			if len(midOrder) > 1 {
				left, right := true, true
				if i == 0 {
					// at most left, tree at right
					left = false
				} else if i == len(midOrder)-1 {
					// most right, tree at left
					right = false
				}

				// split
				if left {
					parentNode.left = new(TreeNode)
					constructTreeNodes(preOrder, index+1, parentNode.left, midOrder[:i])
				}
				if right {
					parentNode.right = new(TreeNode)
					constructTreeNodes(preOrder, index+i+1, parentNode.right, midOrder[i+1:])
				}
			}

			// finished
			return
		}
	}

	panic("not suppose to be here")
}
