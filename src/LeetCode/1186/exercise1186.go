package _186

import "math"

/**
 * CreatedBy: wlei at 2020/6/11
 * Description:
 * 1、问题可以转化为：删除一个元素后，求剩余序列的最大连续子序列之和；将所有得到的和的最大值与原序列的最大连续子序列之和比较取最大值；时间复杂度n^2，超时,问题在于每次计算剩余序列的最大连续子序列之和时重复计算了很大一部分的序列和
 * 2、既然重复计算了，那就想到使用动态规划思想，记忆化之前搜索过的结果，建立不同状态下的状态转移方程
 * 3、 1）状态转移方程定义：令 f[n] 为从左往右顺序的且包含最后一个元素的连续子序列的和的最大值， g[n] 为从左往右顺序删除一个元素后并且包含剩余元素的最后一个元素的的和的最大值，注意这里的最大值不是全局子序列的最大值，
 * 		全局子序列的最大值属于 f[n](不删除元素的全局最大值) 或者 g[n]（删除一个元素后的全局最大值）的一个元素，有状态转移方程： f[n] = max{f[n-1] + arr[n], arr[n]]; g[n] = max{f[n-1], g[n - 1] + arr[n]}
 * 		求出所有的g[n],然后遍历得到其中的全局最大值；
 * 	   2）按照1）中的包含最后一个元素的连续子序列和最大值的定义，分别求出从左到右顺序的最大和的集合： frontToEnd[n] 以及从右到左顺序的最大和的集合： endToFront[n]; 那么当删去 第i个元素后， 只需要计算一个有可能成为全局最大和
 * 		的情况： tempRes = frontToEnd[n - 1] + endToFront[n + 1]; 将所有删除一个元素的情况得到的tempRes和不删除元素的最大连续子序列之和比较得到最大值；
 */

func maximumSum(arr []int) int {
	frontToEnd := make([]int, len(arr))
	endToFront := make([]int, len(arr))
	frontToEnd[0] = arr[0]
	endToFront[len(arr)-1] = arr[len(arr)-1]
	res := arr[0]

	for index := 1; index < len(frontToEnd); index++ {
		frontToEnd[index] = frontToEnd[index-1] + arr[index]
		if frontToEnd[index] < arr[index] {
			frontToEnd[index] = arr[index]
		}
		if frontToEnd[index] > res {
			res = frontToEnd[index]
		}

	}

	for index := len(arr) - 2; index >= 0; index-- {
		endToFront[index] = endToFront[index+1] + arr[index]
		if endToFront[index] < arr[index] {
			endToFront[index] = arr[index]
		}
	}

	for index := 1; index < len(arr)-1; index++ {
		temp := frontToEnd[index-1] + endToFront[index+1]
		if temp > res {
			res = temp
		}
	}
	return res
}

func maximumSum1(arr []int) int {
	sums := make([]int, len(arr))
	g := make([]int, len(arr))
	sums[0] = arr[0]
	g[0] = -200001

	for index := 1; index < len(arr); index++ {
		sums[index] = sums[index-1] + arr[index]
		if sums[index] < arr[index] {
			sums[index] = arr[index]
		}
		g[index] = g[index-1] + arr[index]
		if g[index] < sums[index-1] {
			g[index] = sums[index-1]
		}
	}

	res := math.MinInt32
	for index := 0; index < len(arr); index++ {
		if res < sums[index] {
			res = sums[index]
		}

		if res < g[index] {
			res = g[index]
		}

	}
	return res
}
