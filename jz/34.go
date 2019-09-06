package jz

var paths [][]int

// 二叉树中和为某一值的路径
//
// 输入一颗二叉树和一个整数，打印出二叉树中结点值的和为输入整数的所有路径。路径定义为从树的根结点开始往下一直到叶结点所经过的结点形成一条路径。
//
// 使用path来存储当前的路径 如果当前的路径满足条件 则保存到最后的结果集里
// 这里判断当前的值是不是满足条件使用了递减法 一直将所需的value减到0 而且应用值传递 使得在递归的时候 当前修改的值不会影响到外部的值
func findPath(root *TreeNode, val int) [][]int {
	// Init path
	paths = make([][]int, 0)

	find(root, val, make([]int, 0))

	return paths
}

func find(node *TreeNode, val int, path []int) {
	if node == nil {
		return
	}

	// Add current node to the path
	path = append(path, node.value)
	// Calc the val
	val -= node.value
	// Check now val is 0
	if val == 0 {
		// This is a path, does not need to go deeper any further
		paths = append(paths, path)
		// Remove latest added node
		path = path[:len(path)-1]
		return
	} else {
		if node.left == nil && node.right == nil {
			// This is a leaf
			// Remove latest added node
			path = path[:len(path)-1]
			return
		} else {
			// Left tree
			if node.left != nil {
				find(node.left, val, path)
			}
			// Right tree
			if node.right != nil {
				find(node.right, val, path)
			}
		}
	}
}
