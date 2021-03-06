package built_in

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给定nums构建Tree
func GenTreeByNums(nums []int, nilNum int) *TreeNode {
	if len(nums) == 0 || nums[0] == nilNum {
		return nil
	}
	root := &TreeNode{Val: nums[0]}
	nums = nums[1:]
	queueNode := []*TreeNode{root} // 维护上一层包含空节点的所有节点（父节点）
	for len(nums) != 0 {
		numPreLevelNodes := len(queueNode)
		for idx := 0; idx < numPreLevelNodes; idx++ {
			if 2*idx < len(nums) && nums[2*idx] != nilNum {
				queueNode[idx].Left = &TreeNode{Val: nums[2*idx]}
				queueNode = append(queueNode, queueNode[idx].Left)
			} else {
				queueNode = append(queueNode, nil)
			}
			if 2*idx+1 < len(nums) && nums[2*idx+1] != nilNum {
				queueNode[idx].Right = &TreeNode{Val: nums[2*idx+1]}
				queueNode = append(queueNode, queueNode[idx].Right)
			} else {
				queueNode = append(queueNode, nil)
			}
		}
		queueNode = queueNode[numPreLevelNodes:]
		nums = nums[Min(numPreLevelNodes*2, len(nums)):]
	}
	return root
}

// 先序遍历(Morris遍历)
func PreorderTraversalMorris(root *TreeNode, visitFunc func(node *TreeNode)) {
	// 1. 如果cur无左子树，cur向右移动（cur=cur.right）
	// 2. 如果cur有左子树，找到cur左子树上最右的节点，记为mostRight
	//    a. 如果mostRight的right指针指向空，让其指向cur，cur向左移动（cur=cur.left）
	//    b. 如果mostRight的right指针指向cur，让其指向空，cur向右移动（cur=cur.right）
	// 3. 重复1、2,直到cur为空

	cur := root
	var mostRight *TreeNode
	for cur != nil {
		// cur表示当前节点，mostRight表示cur的左子树的最右节点
		mostRight = cur.Left
		if mostRight != nil {
			// cur有左子树，找到cur左子树的最右节点
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				// mostRight的right指向空，让其指向cur，cur向左移动
				mostRight.Right = cur
				visitFunc(cur) // 先序中序访问不同点
				cur = cur.Left
			} else {
				// mostRight的right指向cur，让其指向空，cur向右移动
				mostRight.Right = nil
				cur = cur.Right
			}
		} else {
			// 如果cur无左子树，cur向右移动
			visitFunc(cur)
			cur = cur.Right
		}
	}
}

// 中序遍历(Morris遍历)
func InorderTraversalMorris(root *TreeNode, visitFunc func(node *TreeNode)) {
	// 1. 如果cur无左子树，cur向右移动（cur=cur.right）
	// 2. 如果cur有左子树，找到cur左子树上最右的节点，记为mostRight
	//    a. 如果mostRight的right指针指向空，让其指向cur，cur向左移动（cur=cur.left）
	//    b. 如果mostRight的right指针指向cur，让其指向空，cur向右移动（cur=cur.right）
	// 3. 重复1、2,直到cur为空

	cur := root
	var mostRight *TreeNode
	for cur != nil {
		// cur表示当前节点，mostRight表示cur的左子树的最右节点
		mostRight = cur.Left
		if mostRight != nil {
			// cur有左子树，找到cur左子树的最右节点
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				// mostRight的right指向空，让其指向cur，cur向左移动
				mostRight.Right = cur
				cur = cur.Left
			} else {
				// mostRight的right指向cur，让其指向空，cur向右移动
				mostRight.Right = nil
				visitFunc(cur) // 先序中序访问不同点
				cur = cur.Right
			}
		} else {
			// 如果cur无左子树，cur向右移动
			visitFunc(cur)
			cur = cur.Right
		}
	}
}

// 后序遍历(Morris遍历)
func PostorderTraversalMorris(root *TreeNode, visitFunc func(node *TreeNode)) {
	// 1. 如果cur无左子树，cur向右移动（cur=cur.right）
	// 2. 如果cur有左子树，找到cur左子树上最右的节点，记为mostRight
	//    a. 如果mostRight的right指针指向空，让其指向cur，cur向左移动（cur=cur.left）
	//    b. 如果mostRight的right指针指向cur，让其指向空，cur向右移动（cur=cur.right）
	// 3. 重复1、2,直到cur为空

	cur := root
	var mostRight *TreeNode
	for cur != nil {
		// cur表示当前节点，mostRight表示cur的左子树的最右节点
		mostRight = cur.Left
		if mostRight != nil {
			// cur有左子树，找到cur左子树的最右节点
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				// mostRight的right指向空，让其指向cur，cur向左移动
				mostRight.Right = cur
				cur = cur.Left
			} else {
				// mostRight的right指向cur，让其指向空，cur向右移动
				mostRight.Right = nil
				visitEdge(cur.Left, visitFunc) //  # 在第二次到达时，逆序访问cur节点的左子树的右边界
				cur = cur.Right
			}
		} else {
			// 如果cur无左子树，cur向右移动
			cur = cur.Right
		}
	}
	visitEdge(root, visitFunc)
}

// 访问边界
func visitEdge(node *TreeNode, visitFunc func(node *TreeNode)) {
	tail := reverseEdge(node)
	cur := tail
	for cur != nil {
		visitFunc(cur)
		cur = cur.Right
	}
	reverseEdge(node)
}

// 反转右边界
func reverseEdge(node *TreeNode) *TreeNode {
	var pre, next *TreeNode
	for node != nil {
		next = node.Right
		node.Right = pre
		pre = node
		node = next
	}
	return pre
}
