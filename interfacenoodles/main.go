package main

import (
	"fmt"
	"shapes"
)

type shape interface {
	area() float64
}

func main() {

	c := shapes.Circle{3}
	r := shapes.Rect{3, 5}
	s := []shape{c, r}
	fmt.Printf("The area of the first shape is %v\n", s[0].area())
	fmt.Printf("The area of the second shape is %v\n", s[1].area())

}
