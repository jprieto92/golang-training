package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	counter := 0
	const goroutines = 100
	var wg sync.WaitGroup
	var m sync.Mutex
	//var m sync.RWMutex // More relaxed memory model
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			m.Lock()
			v := counter
			v++
			//m.Lock() // Lock for rwmutex
			counter = v
			m.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
