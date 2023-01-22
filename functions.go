package main

import "fmt"

func main() {
	//1. multiple return values

	val := func(a, b int) (int, int) {
		// go require explicit returns, it won't automatically return the value of the last expression
		return a, b
	}
	// multiple assignment
	c, d := val(1, 9)
	fmt.Println(c, d)

	//2. accepting a variable number of arguments (Variadic functions)
	sum := func (nums ...int) {
		 //nums becomes []int 
		fmt.Print(nums," ")
		total := 0
		for _, num := range nums {
			total += num
		}
		fmt.Println(total)
	}
	
		sum(1,2)
		sum(1,2,3)
		numslice := []int{1,2,3,4}
		sum(numslice...)
	//3. ability to form closures

}
