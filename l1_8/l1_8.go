package l1_8

import "fmt"

// Функция меняет i-ый (с 1 (по человечески)) элемент. Номер элемента
func MainFunc(num int, bitNum int) {
	var x int = num
	fmt.Printf("Old digit like: %d, old binary like: %b\n ", num, num)
	x ^= 1 << (bitNum - 1)
	fmt.Printf("digit like: %d, binary like: %b\n ", x, x)
}
