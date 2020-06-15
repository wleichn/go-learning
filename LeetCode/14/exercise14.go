package _4

/**
 * CreatedBy: wlei at 2020/6/15
 * Description:
 *
 */

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	str := strs[0]
	for index := 1; index < len(strs); index++ {
		lenMin := len(str)
		if lenMin > len(strs[index]) {
			lenMin = len(strs[index])
		}
		i := 0
		for ; i < lenMin; i++ {
			if str[i] != strs[index][i] {
				break
			}
		}
		str = str[0:i]
	}
	return str
}
