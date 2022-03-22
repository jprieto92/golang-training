package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	c := make(chan int)
	fmt.Println("ROUTINES", runtime.NumGoroutine())

	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(m int) {
				for j := 0; j < 10; j++ {
					c <- j*10 + m
				}
				wg.Done()
			}(i)
			fmt.Println("ROUTINES", runtime.NumGoroutine())
		}
		wg.Wait()
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("ROUTINES", runtime.NumGoroutine())
}
