package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	(*p).first = "Julia"
	(*p).last = "Gonzalez"
	(*p).age++
}

func main() {
	p1 := person{
		first: "Julio",
		last:  "Perez",
		age:   30,
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
