package main

import (
	"fmt"
)

func main() {
	//1. if else
	if 7%2 == 0 {
		fmt.Println("7 is even")

	} else {
		fmt.Println("7 is odd")
	}
	//2. if without else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	//3. if / else if /else
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	//4. There is no ternary if in go

}
