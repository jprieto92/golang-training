package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() {
	fmt.Println("Hi, my name is ", p.first, p.last, "and i am ", p.age, "years old.")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	// Correct
	p := person{
		first: "Javier",
		last:  "Prieto",
		age:   30,
	}
	saySomething(&p)

	// Incorrect - Do not compile
	//saySomething(p)
}
