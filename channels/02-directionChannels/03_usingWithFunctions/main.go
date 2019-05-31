package main

import "fmt"

func main() {

	c := make(chan int)
	// Send
	go foo(c)

	// Recieve
	bar(c)

	fmt.Println("about to exit")
}

// Send

func foo(c chan<- int) {
	c <- 42
}

// Recieve
func bar(c <-chan int) {
	fmt.Println(<-c)
}
