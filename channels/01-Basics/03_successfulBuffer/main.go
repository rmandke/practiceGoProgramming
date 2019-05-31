// This code works because we have defined a buffer for the channel. The buffer allows the channel to store values even when the channel is locked.
package main

import "fmt"

func main() {

	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)

}
