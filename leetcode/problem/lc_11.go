package prob

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	maxC := 0.0
	for i := 0; i < len(height); i++ {
		for j := i; j < len(height); j++ {
			area := float64(j-i) * math.Min(float64(height[i]), float64(height[j]))
			if area > maxC {
				maxC = area
			}
		}
	}

	return int(maxC)
}

func Test11() {
	fmt.Println(maxArea([]int{1, 7}))
}
