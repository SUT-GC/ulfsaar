package prob

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	r := 0
	PNFlag := 0
	SNFlag := 0
	ENFlag := 0
	OverPIntFlag := 0
	OverNIntFlag := 0

	for _, c := range str {
		if c == 43 && PNFlag == 0 && SNFlag == 0 {
			PNFlag = 1
			continue
		}

		if c == 45 && PNFlag == 0  && SNFlag == 0{
			PNFlag = 2
			continue
		}

		if c == 32 && PNFlag == 0 && SNFlag == 0 {
			continue
		}

		if c >= 48 && c <= 57 {
			if SNFlag == 0 && ENFlag == 0 {
				SNFlag = 1
			}

			if ENFlag == 1 {
				continue
			}

			r = r*10 + int(c-48)
			if PNFlag <= 1 && r > math.MaxInt32 {
				OverPIntFlag = 1
				break
			}

			if PNFlag > 1 && -r < math.MinInt32 {
				OverNIntFlag = 1
				break
			}

		} else if ENFlag == 0 {
			ENFlag = 1
		}

		if SNFlag == 0 && c != 43 && c != 45 {
			break
		}
	}

	if PNFlag > 1 {
		r = -r
	}

	if OverPIntFlag == 1 {
		return math.MaxInt32
	}

	if OverNIntFlag == 1 {
		return math.MinInt32
	}

	return r
}

func Test7() {
	fmt.Println(myAtoi("0-1"))
}
