package main

import (
	"fmt"
	"time"
)

func main() {
	//1. basic switch
	i := 5
	fmt.Printf("write %d as ", i)
	switch i {
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	case 6:
		fmt.Println("Six")
	}
	// multiple expressions in the same case statement.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend.")
	default:
		fmt.Println("It's a weekday.")
	}
	//switch without an expression (work as if else statement)
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	// type switch compares types instead of values.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println(("bool"))
		case int:
			fmt.Println("int")
		default:
			fmt.Printf("don't know type %T \n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("meow")

}
