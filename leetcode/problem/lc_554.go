package prob

import (
	"fmt"
)

func findMinAndMoreC(a []int) (int, int) {
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}

	moreC := 0
	for _, v := range a {
		if v > min {
			moreC++
		}
	}

	return min, moreC
}

func sum(a []int) int {
	r := 0

	for _, v := range a {
		r += v
	}

	return r
}
func leastBricks(wall [][]int) int {
	temp := make([]int, len(wall))
	complateC := make([]int, len(wall))

	for i, _ := range wall {
		temp[i] = wall[i][0]
	}

	_, count := findMinAndMoreC(temp)
	result := count

	i := 1
	for ; sum(complateC) < len(wall); {
		for j := 0; j < len(wall); j++ {
			if i >= len(wall[j]) {
				complateC[j] = 1
			}
			if complateC[j] == 0 {
				temp[j] += wall[j][i]
			} else {
				temp[j] += 0
			}

			_, count = findMinAndMoreC(temp)
			if count < result {
				result = count
			}
		}
		i++
	}

	return result
}

func Test554() {
	fmt.Println(leastBricks([][]int{
		[]int{1, 2, 2, 1},
		[]int{3, 1, 2},
		[]int{1, 3, 2},
		[]int{2, 4},
		[]int{3, 1, 2},
		[]int{1, 3, 1, 1},
	}))
}
