package jz

type TreeLinkNode struct {
	value  int
	left   *TreeLinkNode
	right  *TreeLinkNode
	parent *TreeLinkNode
}

// 二叉树的下一个结点
//
// 给定一个二叉树和其中的一个结点，请找出中序遍历顺序的下一个结点并且返回。注意，树中的结点不仅包含左右子结点，同时包含指向父结点的指针。
//
// 如果一个节点的右子树不为空，那么该节点的下一个节点是右子树的最左节点；否则，向上找第一个左链接指向的树包含该节点的祖先节点。
//
// `node` the node in the tree
func GetNext(node *TreeLinkNode) *TreeLinkNode {
	if node.right != nil {
		// has right sub tree
		return mostLeftNode(node.right)
	} else {
		// does not has the right sub tree, go to parent's right tree
		for {
			if node.parent.right != node {
				// we found the right tree
				// find most left node
				return mostLeftNode(node.parent.right)
			} else {
				// keep going up
				node = node.parent
			}
		}
	}
}

func mostLeftNode(node *TreeLinkNode) *TreeLinkNode {
	for {
		if node.left != nil {
			node = node.left
		} else {
			// this is the most left node
			break
		}
	}

	return node
}
