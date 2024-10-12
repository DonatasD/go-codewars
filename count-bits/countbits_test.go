package countbits

import (
	"fmt"
	"testing"
)

func TestCountBits(t *testing.T) {

	tests := []struct {
		n        uint
		expected int
	}{
		{n: 0, expected: 0},
		{n: 10, expected: 2},
		{n: 13, expected: 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("CountBits(%v)", tt.n), func(t *testing.T) {
			if result := CountBits(tt.n); result != tt.expected {
				t.Errorf("CountBits for(%v):	Actual %v Should be %v", tt.n, result, tt.expected)
			}
		})
	}
}

func TestCountBitsV2(t *testing.T) {

	tests := []struct {
		n        uint
		expected int
	}{
		{n: 0, expected: 0},
		{n: 10, expected: 2},
		{n: 13, expected: 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("CountBits(%v)", tt.n), func(t *testing.T) {
			if result := CountBits(tt.n); result != tt.expected {
				t.Errorf("CountBits for(%v):	Actual %v Should be %v", tt.n, result, tt.expected)
			}
		})
	}
}
