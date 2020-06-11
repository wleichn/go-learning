package exercise9

import "testing"

/**
 * CreatedBy: wlei at 2020/6/10
 * Description:
 *
 */

func TestIsPalindrome(t *testing.T) {
	samples := []int{121, -121}
	results := []bool{true, false}

	for index, sample := range samples {
		if results[index] != isPalindrome(sample) {
			t.Errorf("Error result for sample %d", index)
		}
	}
}
