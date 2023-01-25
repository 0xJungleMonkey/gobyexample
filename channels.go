//channels: a complement to goroutines in concurrent Go programs
//channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine 
// and receive those values into another goroutine.

package main
import "fmt"
func main(){
	//create a new channel with make(chan val-type). Channels are typed by the values they convey
	messages := make(chan string)
	//send a value into a channel using the channel <-syntax from a new goroutine
	go func() { messages <- "ping" }()
	// the <-channel syntax receives a value from the channel.
	msg := <-messages
	fmt.Println(msg)
	
}
// By default sends and recieves block until both the sender and receiver are ready. 
// This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.
