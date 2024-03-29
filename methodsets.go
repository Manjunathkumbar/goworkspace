package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

//pointer receiver
func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{5}
	info(&c)
}

/*a NON-POINTER RECEIVER
works with values that are POINTERS or NON-POINTERS.
a POINTER RECEIVER
only works with values that are POINTERS.
*/
