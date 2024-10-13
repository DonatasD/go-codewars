package sumdigpower

import (
	"fmt"
	"testing"
)

func TestSumDigPow(t *testing.T) {

	tests := []struct {
		a        uint64
		b        uint64
		expected []uint64
	}{
		{a: 1, b: 10, expected: []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{a: 1, b: 100, expected: []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 89}},
		{a: 10, b: 89, expected: []uint64{89}},
		{a: 10, b: 100, expected: []uint64{89}},
		{a: 90, b: 100, expected: nil},
		{a: 89, b: 135, expected: []uint64{89, 135}},
		{a: 12157692622039623479, b: 12157692622039625652, expected: []uint64{12157692622039623539}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("SumDigPow(%v, %v)", tt.a, tt.b), func(t *testing.T) {
			if result := SumDigPow(tt.a, tt.b); len(result) != len(tt.expected) {
				t.Errorf("SumDigPow for(%v, %v):	Actual %v Should be %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
