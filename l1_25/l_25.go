package l1_25

import (
	"fmt"
	"time"
)

// Реализация своего таймера

func MainFuncTimer() <-chan int {
	var ch = make(chan int)
	ticker := time.NewTicker(time.Second * 2)
	closer := time.After(time.Second * 13)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("doing something after ticker")
			case <-closer:
				ch <- 1
			}
		}
	}()
	return ch
}
