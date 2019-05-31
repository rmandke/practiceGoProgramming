// This code works because there are multiple go routines. Although the main routine is locked, the other go routine is running and is able to recieve the variable.
package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

}
