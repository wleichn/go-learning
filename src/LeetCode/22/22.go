package _2

/**
 * CreatedBy: wlei at 2020/6/9
 * Description:
 *
 * 搜索与回溯： 利用左括号剩余放置数目永远小于等于右括号的剩余放置数目的机制，递归搜索，并回溯
 * 可以考虑动态规划， 但是没有想到适合实现的递归方程
 *
 * 注意在go语言中参数传递时只有值传递，但是对于切片来说，由于切片是一个结构体，保存了数组的地址，实际上是可以改变传入切片在函数外部的值，但是不能增加函数外部的切片的长度，
 * 函数内部执行append后产生的已经是一个新的slice对象，虽然该对象中的数组指针起始地址没有变化
 */

func generateParenthesis(n int) []string {
	var results []string
	leftBracketNum := n
	rightBracketNum := n

	results = generateBrackets(leftBracketNum, rightBracketNum, "", results)
	return results
}

func generateBrackets(leftBracketNum, rightBracketNum int, currentStatus string, results []string) []string {
	if rightBracketNum == 0 {
		results = append(results, currentStatus)
		return results
	}

	if leftBracketNum < rightBracketNum {
		if leftBracketNum != 0 {
			leftBracketNum--
			currentStatus += "("
			results = generateBrackets(leftBracketNum, rightBracketNum, currentStatus, results)
			leftBracketNum++
			currentStatus = currentStatus[0 : len(currentStatus)-1]
		}

		rightBracketNum--
		currentStatus += ")"
		results = generateBrackets(leftBracketNum, rightBracketNum, currentStatus, results)
		rightBracketNum++
	}

	if leftBracketNum == rightBracketNum {
		leftBracketNum--
		currentStatus += "("
		results = generateBrackets(leftBracketNum, rightBracketNum, currentStatus, results)
	}
	return results
}
