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

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range m {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("******")
	}
}
