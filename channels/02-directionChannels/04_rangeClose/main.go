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

	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

// Recieve
func bar(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}

}
