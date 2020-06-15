package _4

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	samples := [][]string{
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
	}
	results := []string{"fl", ""}

	for index, sample := range samples {
		result := longestCommonPrefix(sample)
		if result != results[index] {
			t.Errorf("error result %s for sample %d", result, index)
		}
	}
}
