package l1_6

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func WriterToChannel(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- rand.Intn(100)
	}
	close(ch)
}

/**
Реализация всех видов остановки горутин
*/

func MagerFunc() {
	ch := make(chan int, 10)
	go WriterToChannel(ch)

	// Запуск функций с реализацией всех остановок
	var (
		shutDownChannel = make(chan struct{})
		wg              sync.WaitGroup
	)

	// Использование канала уведомлений
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer log.Println("shutDownChannel closed")

		for {
			select {
			case <-shutDownChannel:
				return
			case res, ok := <-ch:
				if !ok {
					return
				}
				fmt.Println("res:", res)
			}
		}
	}()

	// Закрытие горутины через context
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()

		ticker := time.NewTicker(time.Millisecond * 1000)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				_, ok := <-ch
				if !ok {
					return
				}
			case <-ctx.Done():
				log.Println("context is used")
				return
			}
		}

	}(ctx)

	// Закрытие канала через системные сигналы
	wg.Add(1)
	go func() {
		defer wg.Done()
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

		<-exit
		log.Println("exit sign was gotten")
	}()

	// Выход из горутины по условию
	wg.Add(1)
	go func(limitParam int) {
		defer wg.Done()

		for val := range ch {
			if val > limitParam {
				log.Println("Condition goroutine is closed")
				return
			}
		}
	}(80)

	// Пример с принудительным закрытием горутины
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer log.Println("Goexit is used")
		time.Sleep(time.Millisecond * 200)
		runtime.Goexit()
	}()

	// Запуск закрытия некоторых горутин
	go func() {
		time.Sleep(time.Millisecond * 200)
		// Закрытие канала shutDown
		close(shutDownChannel)
		// Закрытие горутины через контекст функцией отмены
		cancel()
	}()

	wg.Wait()
}
