package main

import "fmt"

func main() {
	values := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range values {
		fmt.Printf("Value in position %d: %d\n", i, v)
	}
	fmt.Printf("%T\n", values)
}
