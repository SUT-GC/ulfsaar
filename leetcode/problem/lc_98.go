package prob

import "fmt"

func midShow(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}

	result := make([]int, 0)
	result = append(result, midShow(root.Left)...)
	result = append(result, root.Val)
	result = append(result, midShow(root.Right)...)

	return result
}

func isValidBST(root *TreeNode) bool {
	l := midShow(root)
	if len(l) == 0 {
		return true
	}

	min := l[0]
	for _, v := range l[1:] {
		if v <= min {
			return false
		}
		min = v
	}

	return true
}

func Test98() {
	//fmt.Println(isValidBST(&TreeNode{5, &TreeNode{1, nil, nil}, &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{6, nil, nil}}}))

	fmt.Println(isValidBST(&TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}))
}
