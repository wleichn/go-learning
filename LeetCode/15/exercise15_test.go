package _5

import (
	"reflect"
	"testing"
)

/**
 * CreatedBy: wlei at 2020/6/12
 * Description:
 *
 */
func TestThreeSum(t *testing.T) {
	samples := [][]int{
		{-2, 0, 1, 1, 2},
		{-1, 0, 1, 2, -1, -4},
		{0, 0, 0},
	}
	results := [][][]int{
		{{-2, 0, 2}, {-2, 1, 1}},
		{{-1, 0, 1},
			{-1, -1, 2}},
		{{0, 0, 0}},
	}

	for index, sample := range samples {
		result := threeSum(sample)
		if !reflect.DeepEqual(result, results[index]) {
			t.Errorf("Error result %v for sample %d", result, index)
		}
	}
}
