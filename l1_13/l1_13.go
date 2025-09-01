package l1_13

import "fmt"

/*
	Смена чисел между переменными несколькими способами
*/

func ReplaceBothVariables() {
	var1, var2 := 10, 20
	fmt.Printf("first variable is %d; second variable is %d\n", var1, var2)

	// Смена по умолчанию языка Go
	var1, var2 = var2, var1
	fmt.Printf("first variable is %d; second variable is %d\n", var1, var2)

	// Использование арифметики
	var1 = var1 + var2
	var2 = var1 - var2
	var1 = var1 - var2
	fmt.Printf("AFTER MATHMATIC OPERATIONS: first variable is %d; second variable is %d\n", var1, var2)

	// Использование xor
	var1 = var1 ^ var2
	var2 = var1 ^ var2
	var1 = var1 ^ var2
	fmt.Printf("AFTER MATHMATIC OPERATIONS: first variable is %d; second variable is %d\n", var1, var2)
}
