package _9

import "math"

/**
 * CreatedBy: wlei at 2020/6/10
 * Description:
 *
 */

func divide(dividend int, divisor int) int {
	sign := dividend / math.Abs(dividend) * divisor
}
