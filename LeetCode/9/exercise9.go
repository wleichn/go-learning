package exercise9

/**
 * CreatedBy: wlei at 2020/6/10
 * Description:
 *
 */

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	raw := x
	new := 0
	for x > 0 {
		new = new*10 + x%10
		x = x / 10
	}

	return new == raw
}
