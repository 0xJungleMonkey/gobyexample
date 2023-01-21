package main

import (
	"fmt"
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

}
