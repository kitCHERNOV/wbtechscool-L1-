package l1_9

import (
	"fmt"
)

// Создание конвеера, состоящего из двух каналов и двух нод
// I. Нода выполняет запись в канал значений из переданного массива и формирует канал
func firstNode(arr []int) <-chan int {
	var firstChan chan int = make(chan int, 10)
	go func() {
		defer close(firstChan)
		for _, v := range arr {
			firstChan <- v
		}
	}()
	return firstChan
}

// II. Нода умножает на 2 число из первичного канала и формирует канал с числами, умноженными на 2
func lastNode(intCh <-chan int) <-chan int {
	var lastChan chan int = make(chan int, 10)
	go func() {
		defer close(lastChan)
		for v := range intCh {
			lastChan <- v * 2
		}
	}()
	return lastChan
}

func MainFunc() {
	myArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	firstChan := firstNode(myArr)
	lastChan := lastNode(firstChan)

	for doubledValue := range lastChan {
		fmt.Println(doubledValue)
	}
}
