package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type react struct {
	width, hight float64
}

type circle struct {
	radius float64
}

func (r react) area() float64 {
	return r.width * r.width
}

func (r react) perim() float64 {
	return 2*r.width + 2*r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("พื้นที่ =", g.area())
	fmt.Println("ปริมาณ =", g.perim())
}

func main() {
	r := react{width: 5, hight: 10}
	c := circle{radius: 15}

	measure(r)
	measure(c)
}
