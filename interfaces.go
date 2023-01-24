// interfaces are named collections of method signatures.
package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}
type triangle struct {
	line1, line2, line3 float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func (t triangle) area() float64 {
	return (t.line1 + t.line2) / 2 * t.line3

}
func (t triangle) perim() float64 {
	return t.line1 + t.line2 + t.line3
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())

}
func main() {
 rect1 := rect{3, 4}
 circle1 := circle{2}
triangle1 := triangle{1, 2, 3}
measure(rect1)
measure(circle1)
measure(triangle1)

}
