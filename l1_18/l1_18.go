package l1_18

import (
	"sync"
)

// Реализация инкрементирующегося счетка из нескольких гортин
type Counter struct {
	mu      sync.Mutex
	counter int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.counter++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	c.mu.Lock()
	value := c.counter
	c.mu.Unlock()
	return value
}

func TestIncrementImplemetation() int {
	var wg sync.WaitGroup

	counter := Counter{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}
	wg.Wait()
	return counter.Value()
}
