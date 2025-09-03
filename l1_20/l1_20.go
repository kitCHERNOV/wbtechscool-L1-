package l1_20

import "fmt"

func ReverseWordsOrder(str string) string {
	newStr := []rune(str)
	n := len(newStr)

	// Переворот строки
	start, end := 0, n-1
	for start < end {
		newStr[start], newStr[end] = newStr[end], newStr[start]
		start++
		end--
	}
	fmt.Println(string(newStr))

	// Переворот каждого слова
	start = 0
	for i := 0; i <= n; i++ {
		if i == n || newStr[i] == ' ' {
			end = i
			for start < end {
				newStr[start], newStr[end-1] = newStr[end-1], newStr[start]
				start++
				end--
			}
			start = i + 1
		}
	}
	return string(newStr)
}
