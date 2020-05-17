/*
Composition of interface:
	type rectangle interface {
		area() float64
		parameter() float64
	}

	type circle interface {
		area() float64
		parameter() float64
	}

	type geometry interface {
		rectangle
		circle
	}

Empty interface
	- var myObj interface = emp{name: "aa"} // is consider as empty interface.
	- To call perticular method we need to typecast as method of this interface not known
	- It used when we have lots of method and decide their relation latter.
*/

package main

import (
	"fmt"
	"math"
)

///////////////////////////////////////
type rectangle interface {
	area() float64
	parameter() float64
}

type circle interface {
	area() float64
	parameter() float64
}

type geometry interface {
	rectangle
	circle
}

////////////////////////////////////////

type rect struct {
	length float64
	width  float64
}

type cir struct {
	redius float64
}

////////////////////////////////////////

func (r *rect) area() float64 {
	return r.length * r.width
}

func (r *rect) parameter() float64 {
	return 2 * (r.length + r.width)
}

func (c *cir) area() float64 {
	return math.Pi * c.redius * c.redius
}

func (c *cir) parameter() float64 {
	return 2 * math.Pi * c.redius
}

//////////////////////////////////////

func shape() {
	var g geometry = &rect{length: 10, width: 5}
	fmt.Println(g, g.area())
	fmt.Println(g, g.parameter())

	var geo interface{} = &cir{redius: 5}
	fmt.Println(geo, geo.(geometry).area())
	fmt.Println(geo, geo.(geometry).parameter())

	switch geo.(type) {
	case int:
		fmt.Println("int")
	case geometry:
		fmt.Println("interface")
	}
}

/*
func main() {
	shape()
}
*/
