/*
Interfaces: Function
Struct: Variables
*/

package main

import (
	"fmt"
	"math"
)

// Geometry functions like area parameter
type Geometry interface {
	area() float64
	parameter() float64
}

// Rectangle length and width
type Rectangle struct {
	length float64
	width  float64
}

// Circle redius
type Circle struct {
	redius float64
}

func (r *Rectangle) area() float64 {
	return r.length * r.width
}

func (r *Rectangle) parameter() float64 {
	return 2 * (r.length + r.width)
}

func (c *Circle) area() float64 {
	return math.Pi * c.redius * c.redius
}

func (c *Circle) parameter() float64 {
	return 2 * math.Pi * c.redius
}

func measure(g Geometry) {
	fmt.Println("Shape: ", g)
	fmt.Println("Area: ", g.area())
	fmt.Println("Parameter: ", g.parameter())
}

// Shape Give Different Geometry for shape
func Shape() {
	rect := Rectangle{20, 10}
	circle := Circle{10}
	measure(&rect)
	measure(&circle)
}

/*
func main() {
	Shape()
}
*/
