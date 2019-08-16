package prob

import (
	"fmt"
	"math"
)

func threeSumClosest(nums []int, target int) int {
	minDiscount := math.MaxInt64
	result := target
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}

			for k := 0; k < len(nums); k++ {
				if k == i || k == j {
					continue
				}

				discount := int(math.Abs(float64(nums[i] + nums[j] + nums[k] - target)))
				if discount < minDiscount {
					minDiscount = discount
					result = nums[i] + nums[j] + nums[k]
				}
			}
		}
	}

	return result
}

func Test16() {
	fmt.Println(threeSumClosest([]int{1, 1, 1, 0}, -100))
}
