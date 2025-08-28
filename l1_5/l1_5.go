package l1_5

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	ch = make(chan int, 5)
	wg sync.WaitGroup
)

func WriterToChannel(ch chan int, timeLimit time.Duration) {
	defer wg.Done()
	timer := time.After(timeLimit)
	for {
		select {
		case ch <- rand.Intn(100):
			continue
		case <-timer:
			close(ch)
			return
		}
	}
}

func ReaderFromChannel(ch chan int) {
	defer wg.Done()

	for res := range ch {
		fmt.Println(res)
	}
}

func Manager(time time.Duration) {

	wg.Add(2)
	go WriterToChannel(ch, time)
	go ReaderFromChannel(ch)

	wg.Wait()
}
