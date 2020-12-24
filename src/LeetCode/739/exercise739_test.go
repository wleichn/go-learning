package _39

import (
	"reflect"
	"testing"
)

/**
 * CreatedBy: wlei at 2020/6/11
 * Description:
 *
 */
func TestDailyTemperatures(t *testing.T) {
	samples := [][]int{
		{73, 74, 75, 71, 69, 72, 76, 73},
	}
	results := [][]int{
		{1, 1, 4, 2, 1, 1, 0, 0},
	}

	for index, sample := range samples {
		result := dailyTemperatures(sample)
		if !reflect.DeepEqual(results[index], result) {
			t.Errorf("Error result %v for sample %d", result, index)
		}
	}
}
