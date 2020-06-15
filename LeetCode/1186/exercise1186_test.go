package _186

import "testing"

/**
 * CreatedBy: wlei at 2020/6/11
 * Description:
 *
 */

func TestMaximumSum(t *testing.T) {
	samples := [][]int{
		{-1, -1, -1, -1},
		{-50},
		{-7, 6, 1, 2, 1, 4, -1},
		{1, -4, -5, -2, 5, 0, -1, 2},
		{8, -1, 6, -7, -4, 5, -4, 7, -6},
	}
	results := []int{-1, -50, 14, 7, 17}
	//results1 := []int{-1, -50, 14, 6, 14}

	maximumSum(samples[3])

	for index, sample := range samples {
		if results[index] != maximumSum(sample) {
			t.Errorf("Error result %d  for sample %d", maximumSum(sample), index)
		}

		if results[index] != maximumSum1(sample) {
			t.Errorf("1: Error resutl %v for sample %d", maximumSum1(sample), index)
		}
	}
}
