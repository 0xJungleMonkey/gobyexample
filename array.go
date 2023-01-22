// Array is a numbered sequence of elements of a specific length. 
// In typical Go code, slices are much more common;
// arrays are useful in some typical scenarios.
package main
import "fmt"
func main(){
	var a [5]int
	//1. create an empty array that will hold exactly 5 ints. The type if elements and length are both part of the array's type. 
	fmt.Println("emp:",a)

	//set a value at an index using the array[index] = value syntax, and get a value with array[index]
	a[4]=100
	fmt.Println("set",a)
	fmt.Println("Get", a[4])
	fmt.Println("len:",len(a))

	//2. declare and initialize an array in one line 
	b := [5]int{1,2,3,4,5}
	fmt.Println("dcl:",b)

	//3. array type are one dimensional, but you are compose types to build multi-dimensional data structures.
	var twoD [2][3]int
for i := 0; i<2; i++ {
	for j :=0; j<3; j++{
		twoD[i][j] = i + j
	}
}
fmt.Println("2d:", twoD)
}