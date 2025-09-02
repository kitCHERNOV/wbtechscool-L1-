package l1_14

import "strings"

// Орпеделение в рантайме типа переданной переменной
func TypeDeterminer(yourValue any) any {
	switch yourValue.(type) {
	case int:
		return yourValue.(int)
	case string:
		return yourValue.(string)
	case bool:
		return yourValue.(bool)
	case chan int:
		return yourValue.(chan int)
	case chan string:
		return yourValue.(chan string)
	case chan bool:
		return yourValue.(chan bool)
	default:
		return yourValue
	}
}

var justString string

func Other() {
	str := "hello world"
	justString = str[:5]
	justString = strings.Clone(str[:5])
}
