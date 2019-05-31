// This program doesn't work as c is a recieve only channel. ie. we can only recieve values from it but cannot send values to it.
package main

import "fmt"

func main() {

	c := make(<-chan int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("------")
	fmt.Printf("%T\n", c)

}
