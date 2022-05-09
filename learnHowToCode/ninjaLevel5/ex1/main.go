package main

import "fmt"

type person struct {
	firstName  string
	lastName   string
	favFlavors []string
}

func main() {
	p1 := person{
		firstName:  "Jose",
		lastName:   "Martinez",
		favFlavors: []string{"Oreo", "Cookies"},
	}
	p2 := person{
		firstName:  "Antonio",
		lastName:   "Hernandez",
		favFlavors: []string{"Nata", "Vainilla", "Fresa"},
	}

	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	for i, v := range p1.favFlavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.firstName)
	fmt.Println(p2.lastName)
	for i, v := range p2.favFlavors {
		fmt.Println(i, v)
	}
}
