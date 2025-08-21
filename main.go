package main

import (
	"fmt"
	"runtime"
	l1_1 "test/l1_1"
	"test/l2_2"
)

func main() {
	// ===== l1.1 ==== //
	fmt.Println("======= L1.1 =======")
	//
	var instance = l1_1.Action{Human: l1_1.Human{"Jessy", 20}}
	instance.IntroduceYourself()
	instance.Talk("hi")
	// ==== end ==== //

	// ===== l1.2 ==== //
	fmt.Println("L1.2")
	runtime.GOMAXPROCS(1)
	arr := []int{2, 4, 6, 8, 10}
	squaredArr := l2_2.ConcurrencyCalculate(arr)
	fmt.Println(squaredArr)
}
