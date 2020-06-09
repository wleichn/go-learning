package _2

import (
	"testing"
)

/**
 * CreatedBy: wlei at 2020/6/9
 * Description:
 *
 * 注意： 切片比较时，循环比较元素比使用reflect中的deepEqual函数要快
 */

func TestGenerateParenthesis(t *testing.T) {
	samples := []int{3}
	resultsForSam := [][]string{
		{"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()"},
	}
	for index, sample := range samples {
		resultsForOut := generateParenthesis(sample)
		if !stringSliceEqual(resultsForOut, resultsForSam[index]) {
			t.Errorf("Error result for sample %d, out: \n %v", index, resultsForOut)
		}
	}

}

func stringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for index, value := range a {
		if value != b[index] {
			return false
		}
	}

	return true
}
