package main

import "fmt"

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func main() {
	x := foo(1, 2, 3, 4, 5, 6)
	fmt.Printf("Sum: %d\n", x)
	y := bar([]int{1, 2, 3, 4, 5, 6})
	fmt.Printf("Sum: %d\n", y)
}
