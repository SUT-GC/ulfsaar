package prob

import "fmt"

func generateTreesMoreNum(points []int) int {
	if len(points) == 1 {
		return 1
	}

	if len(points) == 0 {
		return 0
	}

	result := 0
	for i, _ := range points {
		leftResult := generateTreesMoreNum(points[:i])
		rightResult := generateTreesMoreNum(points[i+1:])

		if leftResult == 0 {
			result += rightResult
		} else if rightResult == 0 {
			result += leftResult
		} else {
			result += leftResult * rightResult
		}
	}

	return result
}

func numTrees(n int) int {
	points := make([]int, 0)
	for i := 1; i <= n; i++ {
		points = append(points, i)
	}

	return generateTreesMoreNum(points)
}

func Test96() {
	r := numTrees(3)
	fmt.Println(r)
}
