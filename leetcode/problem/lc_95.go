package prob

import "fmt"

func generateTreesMore(points []int) []*TreeNode {
	result := make([]*TreeNode, 0)
	if len(points) == 1 {
		result = append(result, &TreeNode{points[0], nil, nil})
		return result
	}

	if len(points) == 0 {
		return result
	}

	for i, _ := range points {
		leftTrees := generateTreesMore(points[:i])
		rightTrees := generateTreesMore(points[i+1:])

		if len(leftTrees) == 0 {
			for _, v := range rightTrees {
				result = append(result, &TreeNode{points[i], nil, v})
			}
		} else if len(rightTrees) == 0 {
			for _, v := range leftTrees {
				result = append(result, &TreeNode{points[i], v, nil})
			}
		} else {
			for _, l := range leftTrees {
				for _, r := range rightTrees {
					result = append(result, &TreeNode{points[i], l, r})
				}
			}
		}
	}

	return result
}

func generateTrees(n int) []*TreeNode {
	points := make([]int, 0)
	for i := 1; i <= n; i++ {
		points = append(points, i)
	}

	return generateTreesMore(points)
}

func Test95() {
	r := generateTrees(3)
	fmt.Println(r)
}
