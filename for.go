package main

import "fmt"

func main() {
	// 1. single condition
	i := 0
	for i < 3 {
		fmt.Println(i)
		i += 1
	}
	// 2. initial/Condition/after for loop
	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}
	// 3. for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function
	for {
		fmt.Println(("loop"))
		break
	}
	// 4. continue
	for n := 0; n <= 7; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
