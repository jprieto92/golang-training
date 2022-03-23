package main

import (
	"fmt"
	"github.com/jprieto92/golang-training/ninjaLevel12/ex2/quote"
	"github.com/jprieto92/golang-training/ninjaLevel12/ex2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
