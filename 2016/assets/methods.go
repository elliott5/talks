// +build OMIT
package main // example adapted from gobyexample.com/methods

import "fmt"

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.width * r.height
}

func (r *rect) double() {
	r.width *= 2
	r.height *= 2
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	rp := &r
	rp.double()
	fmt.Println("area: ", rp.area())
}
