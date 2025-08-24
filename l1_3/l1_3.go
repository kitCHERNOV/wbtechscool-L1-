package l1_3

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func CreateWorkers(numWorkers int) {
	var wg sync.WaitGroup
	// Канал:
	ch := make(chan int, numWorkers) // буфер в n элементов, как и число воркеров
	// Я выбрал это число чтобы при условии более скоростной работы продьюсера данных в канал
	// Можно было бы понимать что воркеры не будут простаивать и на каждого точно найдется кусочек данных
	log.Println(numWorkers)
	// Запись данных в общий канал
	intArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	go func() {
		for _, v := range intArr {
			ch <- v
		}
		close(ch)
	}()

	// Канал завершения работы воркера
	gracefulShutDownChannel := make(chan struct{})

	// worker
	worker := func(sch <-chan struct{}) {
		defer wg.Done()
		for {
			select {
			case <-sch:
				fmt.Println("worker is done")
				return
			case data, isOk := <-ch:
				if !isOk {
					fmt.Println("task is done, worker is closed")
					return
				}
				fmt.Printf("squared number: %d\n", data*data)
			}
		}
	}

	go func() {
		time.Sleep(2 * time.Second)
		close(gracefulShutDownChannel)
	}()

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(gracefulShutDownChannel)
	}

	wg.Wait()
}
