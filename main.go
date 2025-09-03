package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
	"wbtech/l1_1"
	"wbtech/l1_10"
	"wbtech/l1_11"
	"wbtech/l1_12"
	"wbtech/l1_13"
	"wbtech/l1_16"
	"wbtech/l1_17"
	"wbtech/l1_18"
	"wbtech/l1_19"
	"wbtech/l1_2"
	"wbtech/l1_20"
	"wbtech/l1_3" // пакет l1_3 включает решение модуля l1.3 и l1.4
	"wbtech/l1_5"
	"wbtech/l1_6"
	"wbtech/l1_7"
	"wbtech/l1_8"
	"wbtech/l1_9"
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
	case 8:
		l1_8.MainFunc(4, 3)
	case 9:
		l1_9.MainFunc()
	case 10:
		l1_10.MainFunc()
	case 11:
		arr1 := []int{1, 2, 3}
		arr2 := []int{2, 3, 4}
		intersection := l1_11.SetIntersection(arr1, arr2)
		fmt.Printf("Intersection of arr1: %v and arr2: %v is here: %v", arr1, arr2, intersection)
	case 12:
		arr := []string{"cat", "cat", "dog", "cat", "tree"}
		set := l1_12.CreateSet[string](arr)
		fmt.Println(set)
	case 13:
		l1_13.ReplaceBothVariables()
	case 16:
		arr := []int{1, 2, 3, 12, 4, 6, 8, 10, 12, 14, 16, 18}
		l1_16.QuickSort(arr)
		fmt.Println(arr)
	case 17:
		arr := []int{1, 2, 3, 12, 4, 6, 8, 10, 12, 14, 16, 18}
		l1_16.QuickSort(arr)
		target := l1_17.BinarySearch(arr, 12)
		fmt.Println(arr)
		fmt.Println(target)
	case 18:
		incrementedCounter := l1_18.TestIncrementImplemetation()
		fmt.Println(incrementedCounter)
	case 19:
		newRuStr := l1_19.ReverseStringReader("главрыба")
		fmt.Println(newRuStr)
		newEnStr := l1_19.ReverseStringReader("fisherhouse")
		fmt.Println(newEnStr)
		newEmojiStr := l1_19.ReverseStringReader("✅🧩🧵🔒👥⏱️📊⚛️🏁🕵️‍♂️🚀")
		fmt.Println(newEmojiStr)
	case 20:
		newStr := l1_20.ReverseWordsOrder("snow dog sun")
		fmt.Println(newStr)
	}
}
