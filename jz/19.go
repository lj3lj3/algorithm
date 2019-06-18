package jz

var results [][]*bool

// 正则表达式匹配
//
// 请实现一个函数用来匹配包括 '.' 和 '*' 的正则表达式。模式中的字符 '.' 表示任意一个字符，而 '*' 表示它前面的字符可以出现任意次（包含 0 次）。
//
// 在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串 "aaa" 与模式 "a.a" 和 "ab*ac*a" 匹配，但是与 "aa.a" 和 "ab*a" 均不匹配。
//
// 字符串和正则使用独立的位置标识 向后移动 每次匹配查看下一个字符是否是正则 `.`则直接匹配 `*`分为两种情况 0个 则直接跳过`x*`这两个正则 1个及1个以上则直接进行下次匹配 正则位置不动
func RegMatch(str []rune, pattern []rune) bool {
	// init results
	results = make([][]*bool, len(str)+1)
	for key := range results {
		results[key] = make([]*bool, len(pattern)+1)
	}

	return isRegMatch(str, 0, pattern, 0)
}

func isRegMatch(str []rune, strIndex int, pattern []rune, patternIndex int) bool {
	// this results is calculated
	if results[strIndex][patternIndex] != nil {
		return *results[strIndex][patternIndex]
	}

	var result bool
	// check if string is finished and pattern is finished too
	if patternIndex == len(pattern) {
		result = strIndex == len(str)
	} else {
		// check if current index is valid and current position is same or pattern is `.`
		currentMatch := (strIndex < len(str)) && (str[strIndex] == pattern[patternIndex] || pattern[patternIndex] == '.')

		// check if next pattern is `*`
		if patternIndex+1 < len(pattern) && pattern[patternIndex+1] == '*' {
			// split to 0 and more than 0
			// 0 condition, just skip it
			result = isRegMatch(str, strIndex, pattern, patternIndex+2) ||
				(currentMatch && isRegMatch(str, strIndex+1, pattern, patternIndex)) // more than 1 condition, just move the strIndex to the next and keeping the patternIndex
		} else {
			// move to next
			result = currentMatch && isRegMatch(str, strIndex+1, pattern, patternIndex+1)
		}
	}

	// set it back to results
	results[strIndex][patternIndex] = &result

	return result
}
