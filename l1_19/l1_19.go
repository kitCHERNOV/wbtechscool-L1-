package l1_19

func ReverseStringReader(originalString string) string {
	newString := make([]rune, len(originalString))
	for i, v := range originalString {
		newString[len(originalString)-(i+1)] = v
	}
	return string(newString)
}
