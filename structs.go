package main

import "fmt"

type dog struct {
	name string
	age  int
}

func newdog(name string) dog {
	d := dog{name: name}
	d.age = 10
	return d
}

func main() {
	fmt.Println(dog{"Max", 1})
	fmt.Println(dog{name: "Max"})
	fmt.Println(dog{name: "Max", age: 1})
	fmt.Println(&dog{"Max", 1})
	m1 := newdog("Max")
	m2 := &m1
	fmt.Printf("Address of m1=%v:\t%p\n", m1, &m1)
	fmt.Printf("Address of *m2=%v:\t%p\n", *m2, *&m2)

	fmt.Printf("Address of m2=%v:\t%p\n", m2, &m2)

}
