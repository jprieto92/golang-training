package main

import "fmt"

func main() {
	x := "Test"
	switch {
	case x == "Test":
		fmt.Println("case 1")
	case x != "Test":
		fmt.Println("case 2")
	}
}
