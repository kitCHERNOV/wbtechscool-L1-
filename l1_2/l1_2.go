package l1_2

import "sync"

func ConcurrencyCalculate(arr []int) []int {
	var squares = make([]int, len(arr))
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		for i := 0; i < len(arr); i += 2 {
			squares[i] = arr[i] * arr[i]
		}
		wg.Done()
	}()
	go func() {
		for i := 1; i < len(arr); i += 2 {
			squares[i] = arr[i] * arr[i]
		}
		wg.Done()
	}()

	wg.Wait()
	return squares
}
