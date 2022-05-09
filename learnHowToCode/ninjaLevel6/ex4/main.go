package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hi, my name is %s %s and i am %d years old.\n", p.first, p.last, p.age)
}

func main() {
	p1 := person{
		first: "Andrew",
		last:  "Smith",
		age:   32,
	}
	p1.speak()
}
