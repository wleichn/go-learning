package interview_46

/**
 * 动态规划： 基础状态递归方程： dp(str[n]) = dp(str[1:]) + dp(str[2:]) ；当前状态下的数字的分组情况 = 去掉第一位数字的分组数（前一位数字单独替换成字母） + 去掉前两位数字后的分组数（前两位数字结合后替换成字母）
 * 递归过滤规则： 当且仅当 前两位组成的数大于9小于26时（保证这一个数是一个实际的两位数，即非0开头，且可以用字母替换规则替换）才加上dp(str[2:])的情况
 *
 * 注意go语言中字符串分割的使用：str[low:high] ，为前闭后开规则分割
 */

import "strconv"

func translateNum(num int) int {
	numStr := strconv.Itoa(num)

	return findAns(numStr)
}

func findAns(str string) int {
	var ans int
	if len(str) == 1 || len(str) == 0 {
		ans = 1
	}

	if len(str) > 1 {
		ans = findAns(str[1:])
		suffix := str[0:2]
		suffixNum, _ := strconv.Atoi(suffix)
		if suffixNum < 26 && suffixNum > 9 {
			ans += findAns(str[2:])
		}
	}

	return ans
}
