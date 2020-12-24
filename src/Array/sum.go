package Array

func Sum(numbers []int) (sum int) {
	sum = 0
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

/**
1、空格标志符号忽略索引，不加的话number会成为索引而不是值
2、make可以在创建切片时指定需要的长度和容量
*/

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}

/**
slice[low:high]  获取部分切片，如果冒号一侧没有数字则会取到最边缘的元素
*/
