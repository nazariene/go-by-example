package main

import "fmt"

//Go supports methods defined on struct types.
type rect struct {
	width, height int
}

//This area method has a receiver type of *rect.
//In java this is passed-by-reference
func (r *rect) area() int {
	r.width = 1
	return r.width * r.height
}

//Methods can be defined for either pointer or value receiver types.
//This receives a copy of original struct (passed by value)
func (r rect) perim() int {
	r.width = 2
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	//Here we call the 2 methods defined for our struct.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	//Go automatically handles conversion between values and pointers for method calls.
	//You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}