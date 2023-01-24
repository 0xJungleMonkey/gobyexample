// a goroutine is a lightweight thread of execution.
package main
import ("fmt" 
 "time"
)
func f(from string){
	for i :=0; i<3; i++ {
		fmt.Println(from,":",i)
	}
}
func main(){
	// call f in the usual way, running it synchronously.
	 go f("direct")
	// to invoke this function in a goroutine, use go f(s). This new goroutine wil execute concurrently with the callingone.
	// You can also start a goroutine for an anonymous function call.

	 f("goroutine")
	 go func(msg string) {
        fmt.Println(msg)
    }("going")
	time.Sleep(time.Second)
    fmt.Println("done")
}