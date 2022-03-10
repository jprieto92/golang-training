package main

import (
	"fmt"
	"math"
)

type square struct {
	l float64
	w float64
}

type circle struct {
	r float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.l * s.w
}

func (c circle) area() float64 {
	return c.r * math.Pi * 2.0
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	sq := square{
		l: 4,
		w: 5,
	}
	cq := circle{
		r: 4,
	}
	info(sq)
	info(cq)
}
