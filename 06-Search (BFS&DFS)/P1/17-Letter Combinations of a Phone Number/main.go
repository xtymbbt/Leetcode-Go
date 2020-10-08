package _7_Letter_Combinations_of_a_Phone_Number

func letterCombinations(digits string) []string {
	var res = make([]string, 0, 0)
	// 从1开始，string转为整型为49
	for _, digit := range digits {
		switch digit {
		case 50:
			addString(&res, "a", "b", "c")
		case 51:
			addString(&res, "d", "e", "f")
		case 52:
			addString(&res, "g", "h", "i")
		case 53:
			addString(&res, "j", "k", "l")
		case 54:
			addString(&res, "m", "n", "o")
		case 55:
			addString(&res, "p", "q", "r", "s")
		case 56:
			addString(&res, "t", "u", "v")
		case 57:
			addString(&res, "w", "x", "y", "z")
		}
	}
	return res
}

func addString(res *[]string, strings ...string) {
	if len(*res) == 0 {
		*res = append(*res, strings...)
		return
	}
	var cp = make([][]string, len(strings))
	for i, s := range strings {
		cp[i] = make([]string, len(*res))
		copy(cp[i], *res)
		for j := range cp[i] {
			cp[i][j] += s
		}
	}
	*res = make([]string, 0)
	for i := range cp {
		*res = append(*res, cp[i]...)
	}
}