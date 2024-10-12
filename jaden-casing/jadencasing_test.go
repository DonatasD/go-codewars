package jadencasing

import (
	"fmt"
	"testing"
)

func TestJadenCasing(t *testing.T) {

	tests := []struct {
		n        string
		expected string
	}{
		{n: "most trees are blue", expected: "Most Trees Are Blue"},
		{n: "All the rules in this world were made by someone no smarter than you. So make your own.", expected: "All The Rules In This World Were Made By Someone No Smarter Than You. So Make Your Own."},
		{n: "When I die. then you will realize", expected: "When I Die. Then You Will Realize"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("ToJadenCase(%v)", tt.n), func(t *testing.T) {
			if result := ToJadenCase(tt.n); result != tt.expected {
				t.Errorf("ToJadenCase for(%v):	Actual %v Should be %v", tt.n, result, tt.expected)
			}
		})
	}
}

func TestJadenCasingV2(t *testing.T) {

	tests := []struct {
		n        string
		expected string
	}{
		{n: "most trees are blue", expected: "Most Trees Are Blue"},
		{n: "All the rules in this world were made by someone no smarter than you. So make your own.", expected: "All The Rules In This World Were Made By Someone No Smarter Than You. So Make Your Own."},
		{n: "When I die. then you will realize", expected: "When I Die. Then You Will Realize"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("ToJadenCaseV2(%v)", tt.n), func(t *testing.T) {
			if result := ToJadenCase(tt.n); result != tt.expected {
				t.Errorf("ToJadenCaseV2 for(%v):	Actual %v Should be %v", tt.n, result, tt.expected)
			}
		})
	}
}
