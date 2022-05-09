package main

import "fmt"

func foo() int {
	return 4
}

func bar() (int, string) {
	return 5, "hello"
}

func main() {
	x := foo()
	y, z := bar()
	fmt.Printf("%d\n", x)
	fmt.Printf("%d, %s\n", y, z)
}
