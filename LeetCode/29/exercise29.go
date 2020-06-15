package _9

import (
	"math"
)

/**
 * CreatedBy: wlei at 2020/6/10
 * Description:
 * 注意边界值的判定； 同时，为了缩短计算复杂度，每做一次减法，将减数翻倍
 */

func divide(dividend int, divisor int) int {
	var sign = 1
	if dividend < 0 {
		dividend = -dividend
		sign = -1
	}
	if divisor < 0 {
		divisor = -divisor
		if sign > 0 {
			sign = -1
		} else {
			sign = 1
		}
	}

	result := 0
	tempResult := 1
	tempDivisor := divisor
	for dividend >= divisor {
		if result == math.MaxInt32 && sign > 0 {
			return result
		} else if result == math.MaxInt32 && sign < 0 {
			return math.MinInt32
		}

		result = result + tempResult
		dividend -= tempDivisor

		tempDivisor += tempDivisor
		if dividend >= tempDivisor {
			tempResult += tempResult
		} else {
			tempResult = 1
			tempDivisor = divisor
		}
	}
	if sign < 0 {
		result = -result
	}
	return result
}
