package main

import (
	"fmt"
	"sync"
)

func Square(i int) {
	fmt.Println(i * i)
}

// func main() {
// 	val := [...]int{2, 4, 6, 8, 10}
// 	var wg sync.WaitGroup
// 	wg.Add(len(val))
// 	for i := 0; i < len(val); i++ {
// 		go func(i int) {
// 			Square(i)
// 			defer wg.Done()
// 		}(val[i])

// 	}
// 	wg.Wait()
// }

// func main() {
// 	ch := make(chan struct{})
// 	val := [...]int{2, 4, 6, 8, 10}
// 	for i := 0; i < len(val); i++ {
// 		go func(i int) {
// 			Square(i)
// 			ch <- struct{}{} // Уведомляем о завершении
// 		}(val[i])

// 		<-ch // Ждем уведомления о завершении
// 	}
// }

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	val := [...]int{2, 4, 6, 8, 10}
	ready := false
	var wg sync.WaitGroup

	for i := 0; i < len(val); i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			mu.Lock()
			for !ready {
				cond.Wait() // Освобождаем мьютекс и ждем сигнала
			}
			Square(i)
			mu.Unlock()
		}(val[i])

		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			ready = true
			cond.Signal()
			mu.Unlock()
		}()
	}

	wg.Wait()
}
