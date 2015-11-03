package main

import (
	"testing"
)

func ExpectValue(t *testing.T, expected []int, value string, tolerance string) bool {
	result := extractResistorDigits(value, tolerance)
	if result == nil && expected != nil {
		t.Errorf("Unexpected nil for '%s'", value)
		return false
	}
	if len(result) != len(expected) {
		t.Errorf("%s: Expected len %d but got %d", value, len(expected), len(result))
		return false
	}
	for idx, _ := range result {
		if expected[idx] != result[idx] {
			t.Errorf("%s expected[%d] != result[%d] (%d vs. %d)",
				value, idx, idx, expected[idx], result[idx])
		}
	}
	return true
}

func TestExtractResistorValue(t *testing.T) {
	ExpectValue(t, []int{1, 0, 2, 10}, "1k", "5%")
	ExpectValue(t, []int{1, 0, 2, 1}, "1k", "1%")
	ExpectValue(t, []int{1, 0, 2, 1}, "1.0k", "1%")
	ExpectValue(t, []int{1, 0, 0, 1, 1}, "1.00k", "1%")

	ExpectValue(t, []int{1, 0, 3, 10}, "10k", "5%")
	ExpectValue(t, []int{1, 0, 4, 10}, "100k", "5%")
	ExpectValue(t, []int{1, 0, 4, 10}, "100000", "5%")

	ExpectValue(t, []int{2, 3, 7, 2, 10}, "23.7k", "5%")

	ExpectValue(t, []int{1, 5, 10, 10}, "1.5", "5%")
	ExpectValue(t, []int{1, 5, 11, 10}, "0.15", "5%")

	ExpectValue(t, nil, "10k x", "5%") // garbage.
}
