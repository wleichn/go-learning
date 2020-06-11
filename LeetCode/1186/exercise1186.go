package _186

import "math"

/**
 * CreatedBy: wlei at 2020/6/11
 * Description:
 * 1、问题可以转化为：删除一个元素后，求剩余序列的最大连续子序列之和；将所有得到的和的最大值与原序列的最大连续子序列之和比较取最大值；时间复杂度n^2，超时
 * 2、
 */

func maximumSum(arr []int) int {
	return findMaxSumOfSubArr(arr)
}

func findMaxSumOfSubArr(arr []int) int {
	maxSum := math.MinInt32
	tempSum := math.MinInt32
	deleteFlag := false
	indexFlag := 0
	minElement := 0

	for index := 0; index < len(arr); index++ {
		value := arr[index]
		if value < 0 && minElement > value {
			minElement = value
		}
		if tempSum == math.MinInt32 {
			tempSum = value
		} else {
			tempSum += value
		}
		if tempSum > maxSum {
			maxSum = tempSum
		}

		if !deleteFlag && value < 0 && tempSum != value {
			deleteFlag = true
			tempSum -= minElement
			minElement = 0
			indexFlag = index
		}
		if tempSum < 0 {
			tempSum = math.MinInt32
			if deleteFlag {
				index = indexFlag
			}
			deleteFlag = false
			minElement = 0
		}
	}
	if tempSum > 0 && minElement != 0 {
		tempSum -= minElement
		if tempSum > maxSum {
			maxSum = tempSum
		}
	}
	return maxSum
}
