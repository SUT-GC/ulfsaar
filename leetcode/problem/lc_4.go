package prob

import (
	"fmt"
	"math"
)

func sumAndFind(nums1 []int, nums2 []int) float64 {
	sum := make([]int, 0)
	if len(nums1) == 0 {
		sum = nums2
	} else if len(nums2) == 0 {
		sum = nums1
	} else {
		i, j := 0, 0
		for {
			if i == len(nums1) && j == len(nums2) {
				break
			}

			if i == len(nums1) {
				sum = append(sum, nums2[j])
				j++
			} else if j == len(nums2) {
				sum = append(sum, nums1[i])
				i++
			} else {
				if nums1[i] >= nums2[j] {
					sum = append(sum, nums2[j])
					j++
				} else {
					sum = append(sum, nums1[i])
					i++
				}
			}
		}
	}

	if len(sum)%2 == 0 {
		return float64(sum[len(sum)/2]+sum[len(sum)/2-1]) / 2.0
	} else {
		return float64(sum[len(sum)/2])
	}
}

func dichotomizeAndFind(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	sumLength := len(nums1) + len(nums2)
	for i, _ := range nums1 {
		j := sumLength / 2
		if nums2[j] > nums1[i] {
			continue
		}

		if sumLength%2 == 1 {
			return float64(nums2[j+1])
		} else {
			return (math.Max(float64(nums1[i]), float64(nums2[j])) + math.Min(float64(nums1[i+1]), float64(nums2[i+1]))) / 2.0
		}
	}

	return 0.0
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	result := sumAndFind(nums1, nums2)
	//if len(nums1) == 0 || len(nums2) == 0 || nums1[len(nums1)-1] <= nums2[0] || nums2[len(nums2)-1] <= nums1[0] {
	//	result = sumAndFind(nums1, nums2)
	//} else {
	//	result = dichotomizeAndFind(nums1, nums2)
	//}

	return result
}

func Test4() {
	fmt.Println(findMedianSortedArrays([]int{1, 2, 4}, []int{4, 5, 6}))
}
