package _90

/**
并查集： 所有相等的字符使用并查集分类，然后再检查分类结果是否满足不等的情况，不满足则返回false；
并查集合并操作时秩的变化不理解；
全局变量注意初始化
*/

var r [1050]int
var uSet [1050]int

func equationsPossible(equations []string) bool {

	var unEquations []string
	var indexFlag int

	words := make(map[byte]int)

	for i := 0; i < 1050; i++ {
		uSet[i] = i
		r[i] = 0
	}

	for _, str := range equations { // 对 "=="情况进行分类，等于则前后两个字母的祖先相同，同时抽出"!="的情况
		word := str[1]
		if word == '!' {
			unEquations = append(unEquations, str)
		}

		word0 := str[0]
		word1 := str[3]

		class0, exist0 := words[word0]
		class1, exist1 := words[word1]

		if !exist0 {
			words[word0] = indexFlag
			class0 = indexFlag
			indexFlag++
		}

		if !exist1 {
			words[word1] = indexFlag
			class1 = indexFlag
			indexFlag++
		}
		if word != '!' { // 为"=="的情况才合并并查集
			unionJoin(findRoot(class0), findRoot(class1))
		}

	}

	for _, str := range unEquations { // 检查"！="条件是否和分类结果匹配
		word0 := str[0]
		word1 := str[3]

		class0, _ := words[word0]
		class1, _ := words[word1]

		if findRoot(class1) == findRoot(class0) { // 检查两个字符的祖先是否相同
			return false
		}
	}
	return true
}
func unionJoin(i0, i1 int) {
	if i0 == i1 {
		return
	}

	if r[i0] > r[i1] {
		uSet[i1] = i0
	} else {
		if r[i0] == r[i1] {
			r[i1]++
		}
		uSet[i0] = i1
	}
}

func findRoot(index int) int {
	if index != uSet[index] {
		uSet[index] = findRoot(uSet[index])
	}

	return uSet[index]
}
