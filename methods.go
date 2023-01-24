// go supports methods defined on struct types
// receiver
package main

import "fmt"

type rect struct {
	width, height int
}
// pointer receiver type of *rect
func (r *rect) area() int {
	return r.width * r.height
}
// value receiver type of rect
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}
func main() {
	r := rect{3, 4}
	fmt.Println(r.area())
	fmt.Println(r.perim())
	// go automatically handles conversion between values and pointers for method calls. 
	// You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
	rp := &r
	fmt.Println(rp.area())
	fmt.Println(rp.perim())
}
