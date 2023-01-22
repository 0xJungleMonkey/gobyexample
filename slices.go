// slices are giving a more powerful interface to sequences than arrays.
// unlike arrays, slices are types only by the elements they contain(not the number of elements.)

package main

import "fmt"

func main() {
	//1. create an empty slice with non-zero length, use the builtin make.
	// Here we make a slice of strings of length 3( initially zero-valued).
	s := make([]string, 3)
	fmt.Println(s)
	s[0] = "meow"
	s[1] = "?"
	s[2] = "!"
	fmt.Println(s)
	fmt.Println("Len: ", len(s))

	//2. append can change the length of the slice
	s = append(s, "newelement")
	fmt.Println("after:", s)
	fmt.Println("after:", len(s))

	//declare an empty slice
	s2 := []int{}
	fmt.Println(len(s2))
	s2 = append(s2, 0)
	fmt.Println(len(s2))
	fmt.Println(s2)
	//declare and initialize
	s3 := []int{1, 2}
	fmt.Println(len(s3))
	fmt.Println(s3)
	s3 = append(s3, 0)
	fmt.Println(len(s3))
	fmt.Println(s3)

	//3. Slice can be copy'd
	// create an empty slice of the same length as s and copy into c from s
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(len(c))
	fmt.Println(c)

	//4. slice operator slice[low:high]
	l := s[2:6]
	fmt.Println(len(l))
	fmt.Println(l)
	//5. multi-dimensional slices
	twoD := make([][]int,3)
	for i:=0;i<3;i++{
		innerLen := i+1
		twoD[i]=make([]int,innerLen)
		for j:=0; j<innerLen; j++{
			twoD[i][j] = i + j
		}

	}
	fmt.Println("2d: " ,twoD)

}
