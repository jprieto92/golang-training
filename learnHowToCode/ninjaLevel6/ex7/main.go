package main

import "fmt"

func main() {
	x := func(i int) {
		fmt.Printf("Value: %d\n", i)
	}
	x(5)
}
