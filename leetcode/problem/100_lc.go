package prob

import "fmt"

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil && q != nil {
		return false
	}

	if q == nil && p != nil {
		return false
	}

	if q.Val != p.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func Test100() {
	fmt.Println(isSameTree(&TreeNode{1, &TreeNode{2, nil, nil}, nil}, &TreeNode{1, nil, &TreeNode{2, nil, nil}}))
}
