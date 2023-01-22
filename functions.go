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
	sum := func(nums ...int) {
		//nums becomes []int
		fmt.Print(nums, " ")
		total := 0
		for _, num := range nums {
			total += num
		}
		fmt.Println(total)
	}

	sum(1, 2)
	sum(1, 2, 3)
	numslice := []int{1, 2, 3, 4}
	sum(numslice...)
	//3. ability to form closures (anonymous functions)
	// anonymous functions are useful when you want to define a function inline without having to name it.
	intSeq :=func() func() int {
		i := 0
		return func() int {
			i++
			return i
			
		}
	}
	//
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
    fmt.Println(newInts())
//4.Recursion
	fmt.Println(fact(5))
	var fib func(n int) int
    fib = func(n int) int {
        if n < 2 {
            return n
        }
		return fib(n-1) + fib(n-2)
	}
    fmt.Println(fib(7))
}

//4.Recursion
	//use functions that call themselves from within their own code.
func fact(n int) int {
	if n==0 {
		return 1
	}
	return n * fact(n-1)
}

