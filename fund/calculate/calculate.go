package calculate

import "math"

func CalculatePrincipal(annualized float64, years int, dist float64) float64 {
	percent := 1 + annualized
	sum := 0.0
	for i := 1; i <= years; i++ {
		sum = sum + 12*math.Pow(percent, float64(i))
	}

	return dist / sum
}
