package _9

import "testing"

/**
 * CreatedBy: wlei at 2020/6/15
 * Description:
 *
 */

func TestDivide(t *testing.T) {
	dividends := []int{
		10, 7, -2147483648,
	}
	divisors := []int{
		3, -3, -1,
	}
	results := []int{
		3, -2, 2147483647,
	}
	for index := 0; index < len(dividends); index++ {
		dividend := dividends[index]
		divisor := divisors[index]
		result := divide(dividend, divisor)
		if result != results[index] {
			t.Errorf("error result %d for sample %d", result, index)
		}
	}
}
