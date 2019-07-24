package prob

import (
	"fmt"
	"math"
)

func maxArea2(height []int) int {
	maxC := 0.0
	i := 0
	j := len(height) - 1
	for {
		if i >= j {
			break
		}
		area := float64(j-i) * math.Min(float64(height[i]), float64(height[j]))
		if area > maxC {
			maxC = area
		}

		if height[i] >= height[j] {
			j--
		} else {
			i++
		}
	}
	return int(maxC)
}

func Test11_1() {
	fmt.Println(maxArea2([]int{1,  7}))
}
