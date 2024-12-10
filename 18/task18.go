package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mtx   sync.RWMutex
}

func (c *Counter) Increment() {
	c.mtx.Lock()
	c.value++
	c.mtx.Unlock()
}

func (c *Counter) GetCounter() int {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	numGoroutines := 100
	incrementsPerGoroutine := 1000

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()

	fmt.Printf("Final counter value: %d\n", counter.GetCounter())
}
