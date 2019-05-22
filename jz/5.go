package jz

// 替换空格
//
// 将一个字符串中的空格替换成 "%20"。
//
// Input:
//"A B"
//
//Output:
//"A%20B"
func replaceSpace(str string) string {
	// use slice
	var replacedStr []rune
	for _, char := range str {
		if char == ' ' {
			replacedStr = append(replacedStr, '%', '2', '0')
		} else {
			replacedStr = append(replacedStr, char)
		}
	}

	return string(replacedStr)
}
