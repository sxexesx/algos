package revers_numbers

import "math"

func revNumber(n uint64, dp int) uint64 {
	var res uint64
	for i := 0; i < dp; i++ {
		digit := n % 10
		tens := uint64(math.Pow(10.0, float64(dp-i-1)))
		res += digit * tens
		n = n / 10
	}
	return res
}
