package gimme

import (
	"fmt"
	"testing"
)

func TestGimme(t *testing.T) {

	tests := []struct {
		n        [3]int
		expected int
	}{
		{n: [3]int{2, 3, 1}, expected: 0},
		{n: [3]int{5, 10, 14}, expected: 1},
		{n: [3]int{1, 3, 4}, expected: 1},
		{n: [3]int{15, 10, 14}, expected: 2},
		{n: [3]int{-4, -23, 4}, expected: 0},
		{n: [3]int{-15, -10, 14}, expected: 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Gimme(%v)", tt.n), func(t *testing.T) {
			if result := Gimme(tt.n); result != tt.expected {
				t.Errorf("Gimme for(%v):	Actual %v Should be %v", tt.n, result, tt.expected)
			}
		})
	}
}

func TestGimmeV2(t *testing.T) {

	tests := []struct {
		n        [3]int
		expected int
	}{
		{n: [3]int{2, 3, 1}, expected: 0},
		{n: [3]int{5, 10, 14}, expected: 1},
		{n: [3]int{1, 3, 4}, expected: 1},
		{n: [3]int{15, 10, 14}, expected: 2},
		{n: [3]int{-4, -23, 4}, expected: 0},
		{n: [3]int{-15, -10, 14}, expected: 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("GimmeV2(%v)", tt.n), func(t *testing.T) {
			if result := GimmeV2(tt.n); result != tt.expected {
				t.Errorf("GimmeV2 for(%v):	Actual %v Should be %v", tt.n, result, tt.expected)
			}
		})
	}
}
