// range iterate over elements in a variety of data structures.

package main

import "fmt"

func main() {
	// 1. slice/array
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

	//2. map

	kvs := map[string]string{"a":"apple","b":"banana"}
	// iterate over the pairs of map
	for k,v := range kvs {
		fmt.Printf("%s -> %s\n",k,v)
	}
	//iterate over the keys of a map
	for k := range kvs {
		fmt.Println("Key", k)
	}
	//3. strings 
	//iterates over Unicode code points; The first value is the starting byte index of the rune and the second the rune itself. 
	for i,c := range "go"{
		fmt.Println(i,c)
	}

	

}
