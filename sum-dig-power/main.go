package sumdigpower

import (
	"math"
	"strconv"
)

func SumDigPow(a, b uint64) []uint64 {
	var result []uint64
	for i := a; i <= b; i++ {
		str := strconv.FormatUint(i, 10)
		var count float64
		for j, v := range str {
			count += math.Pow(float64(v-'0'), float64(j+1))
		}
		v, err := strconv.ParseFloat(str, 64)
		if err == nil && count == v {
			v, err := strconv.ParseUint(str, 10, 64)
			if err == nil {
				result = append(result, v)
			}
		}
	}
	return result
}
