package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	data map[int]int
	m    sync.RWMutex
}

func NewConcurrentMap() ConcurrentMap {
	return ConcurrentMap{data: make(map[int]int)}
}

func (c *ConcurrentMap) Add(key, value int, wg *sync.WaitGroup) {
	c.m.Lock()
	c.data[key] = value
	c.m.Unlock()
	wg.Done()
}

func (c *ConcurrentMap) Get(key int) int {
	c.m.RLock() //позволяет нескольким потокам одновременно читать данные из карты, но не изменять и
	defer c.m.RUnlock()
	return c.data[key]
}

func (c *ConcurrentMap) PrintAll() {
	c.m.RLock()
	defer c.m.RUnlock()
	for key, value := range c.data {
		fmt.Println(key, " ", value)
	}
}

func main() {
	cm := NewConcurrentMap()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go cm.Add(i, i*i, &wg)
	}

	wg.Wait()

	// Читаем значения параллельно
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			value := cm.Get(key)
			fmt.Printf("Получено значение для ключа %d: %d\n", key, value)
		}(i)
	}

	wg.Wait()

	fmt.Println("Все значения в ConcurrentMap:")
	cm.PrintAll()
}
