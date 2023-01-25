package main

import (
    "fmt"
    "time"
)
//We can use channels to synchronize execution across goroutines. Here’s an example of using a blocking receive to wait for a goroutine to finish. 
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func main() {

    done := make(chan bool, 1)
    go worker(done)
//Block until we receive a notification from the worker on the channel. 
<-done
//If you removed the <- done line from this program, the program would exit before the worker even started.
}