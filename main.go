package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
	"wbtech/l1_1"
	"wbtech/l1_2"
	"wbtech/l1_3" // пакет l1_3 включает решение модуля l1.3 и l1.4
	"wbtech/l1_5"
	"wbtech/l1_6"
	"wbtech/l1_7"
)

func main() {
	var num int
	fmt.Println("Write a number of task solution")
	fmt.Scanf("%d", &num)

	switch num {
	case 1:
		// ===== l1.1 ==== //
		fmt.Println("======= L1.1 =======")
		var instance = l1_1.Action{Human: l1_1.Human{"Jessy", 20}}
		instance.IntroduceYourself()
		instance.Talk("hi")
		// ==== end ==== //
	case 2:
		// ===== l1.2 ==== //
		fmt.Println("======= L1.2 =======")
		runtime.GOMAXPROCS(1)
		arr := []int{2, 4, 6, 8, 10}
		squaredArr := l1_2.ConcurrencyCalculate(arr)
		fmt.Println(squaredArr)
	case 3:
		// ===== l1.3 ==== //
		fmt.Println("======= L1.3 и L1.4 =======")
		numWorkers := rand.Intn(10) + 1
		l1_3.CreateWorkers(numWorkers)
	case 4:
		// ===== l1.3 ==== //
		fmt.Println("======= L1.3 и L1.4 =======")
		numWorkers := rand.Intn(10) + 1
		l1_3.CreateWorkers(numWorkers)
	case 5:
		// ===== l1.5 ===== //
		l1_5.Manager(time.Second)
	case 6:
		// ===== l1.6 ===== //
		l1_6.MagerFunc()
	case 7:
		l1_7.MainFunc()
	}
}
