package prob

import (
	"fmt"
)

func midShowForNode(root *TreeNode) []*TreeNode {
	nodeResult := make([]*TreeNode, 0)
	if root.Left == nil && root.Right == nil {
		nodeResult = append(nodeResult, root)
		return nodeResult
	}

	if root.Left != nil {
		a := midShowForNode(root.Left)
		nodeResult = append(nodeResult, a...)
	}

	nodeResult = append(nodeResult, root)

	if root.Right != nil {
		a := midShowForNode(root.Right)
		nodeResult = append(nodeResult, a...)
	}

	return nodeResult
}

func recoverTree(root *TreeNode) {
	midShowList := midShowForNode(root)
	var errorNode1 *TreeNode = nil
	var errorNode2 *TreeNode = nil
	for i, _ := range midShowList {
		if i == 0 {
			continue
		}
		if errorNode1 == nil && midShowList[i-1].Val > midShowList[i].Val {
			errorNode1 = midShowList[i-1]
		}

		if errorNode1 != nil && midShowList[i-1].Val > midShowList[i].Val {
			errorNode2 = midShowList[i]
		}
	}

	errorNode1.Val, errorNode2.Val = errorNode2.Val, errorNode1.Val
}

func Test99() {
	tree := TreeNode{1, &TreeNode{3, nil, &TreeNode{2, nil, nil}}, nil}
	recoverTree(&tree)
	fmt.Println(tree)
}
