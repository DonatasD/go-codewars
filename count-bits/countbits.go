package countbits

import "math/bits"

func CountBits(n uint) int {
	count := 0
	for n > 0 {
		count += int(n & 1)
		n >>= 1
	}
	return count
}

func CountBitsV2(n uint) int {
	return bits.OnesCount(n)
}
