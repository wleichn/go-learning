package Array

import (
	"fmt"
)

func Sum(numbers [5]int) (sum int) {
	sum = 0
	for _, number := range numbers {
		sum += number
		fmt.Println(number)
	}
	return
}

/**
空格标志符号忽略索引，不加的话number会成为索引而不是值
*/
