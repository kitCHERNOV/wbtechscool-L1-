package l1_26

import "strings"

func IsStringConsistUniqueLetters(str string) bool {
	str = strings.ToLower(str)
	checkMap := make(map[rune]struct{})
	for _, v := range str {
		if _, ok := checkMap[v]; ok {
			return false
		}
		checkMap[v] = struct{}{}
	}
	return true
}
