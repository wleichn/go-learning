package _39

/**
 * CreatedBy: wlei at 2020/6/11
 * Description:
 * 动态规划思想： 状态转移方程不太好表示，
 * 使用递归搜索实现：从最后一个元素开始往前更新状态，每一个元素的状态通过后一个元素的状态辅助更新：
 * 当前元素如果大于后一个位置的元素，则可以通过后一个位置元素已经更新好的状态跳过不必要的比较
 */

var ans []int

func dailyTemperatures(T []int) []int {
	ans = make([]int, len(T))
	ans[len(ans)-1] = 0
	for index := len(T) - 2; index >= 0; index-- {
		value := T[index]
		ans[index] = findBiggerIndex(value, index, 1, T)
	}

	return ans
}

func findBiggerIndex(value, index, delta int, T []int) int {
	res := 0
	if index+delta >= len(T) {
		return 0
	}

	if value < T[index+delta] {
		res = delta
	} else {
		if ans[index+delta] == 0 {
			return 0
		}
		res = findBiggerIndex(value, index, delta+ans[index+delta], T)
	}
	return res
}
