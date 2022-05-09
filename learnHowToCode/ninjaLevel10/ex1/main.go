package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	d := make(chan int, 1)

	go func() {
		c <- 42
		d <- 50
	}()
	fmt.Println(<-c)
	fmt.Println(<-d)
}
