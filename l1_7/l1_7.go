package l1_7

import (
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

type SafeMap struct {
	sync.RWMutex
	m map[int]int
}

func NewSafeMap(size int) *SafeMap {
	return &SafeMap{m: make(map[int]int, size)}
}

func (sm *SafeMap) WriterToMap() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		number := rand.Intn(100)

		sm.Lock()
		sm.m[i] = number
		sm.Unlock()
	}
}

func MainFunc() {
	safeMAp := NewSafeMap(100)

	wg.Add(3)
	go safeMAp.WriterToMap()
	go safeMAp.WriterToMap()
	go safeMAp.WriterToMap()

	wg.Wait()
}
