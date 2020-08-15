package merge_trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	var result *TreeNode
	if t1 == nil && t2 == nil {
		return nil
	}
	valNode := TreeNode{Val: 0, Left: nil, Right: nil}
	result = &valNode
	transerval(t1, t2, result)
	return result
}

func transerval(t1 *TreeNode, t2 *TreeNode, result *TreeNode) {
	if t1 == nil && t2 == nil {
		return
	}
	switch {
	case t1 == nil:
		result.Val = t2.Val
		if t2.Left != nil {
			valNode := TreeNode{Val: 0, Left: nil, Right: nil}
			result.Left = &valNode
			transerval(nil, t2.Left, result.Left)
		}
		if t2.Right != nil {
			valNode := TreeNode{Val: 0, Left: nil, Right: nil}
			result.Right = &valNode
			transerval(nil, t2.Right, result.Right)
		}
	case t2 == nil:
		result.Val = t1.Val
		if t1.Left != nil {
			valNode := TreeNode{Val: 0, Left: nil, Right: nil}
			result.Left = &valNode
			transerval(t1.Left, nil, result.Left)
		}
		if t1.Right != nil {
			valNode := TreeNode{Val: 0, Left: nil, Right: nil}
			result.Right = &valNode
			transerval(t1.Right, nil, result.Right)
		}
	default:
		result.Val = t1.Val + t2.Val
		if t1.Left != nil || t2.Left != nil {
			setUpNextNode(t1, t2, result, true)
		}
		if t1.Right != nil || t2.Right != nil {
			setUpNextNode(t1, t2, result, false)
		}
	}
}

func setUpNextNode(t1 *TreeNode, t2 *TreeNode, result *TreeNode, isLeft bool) {
	valNode := TreeNode{Val: 0, Left: nil, Right: nil}
	if isLeft {
		result.Left = &valNode
		transerval(t1.Left, t2.Left, result.Left)
	} else {
		result.Right = &valNode
		transerval(t1.Right, t2.Right, result.Right)
	}
}
