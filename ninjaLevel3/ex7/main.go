package main

import "fmt"

func main() {
	x := "kk"

	if x == "Test" {
		fmt.Println(x)
	} else if x == "Test2" {
		fmt.Println(x)
	} else {
		fmt.Println("No test version supported")
	}
}
