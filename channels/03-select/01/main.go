package main

import "fmt"

func main() {

	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)

	receive(eve, odd, quit)

	fmt.Println("about to exit")

}

func receive(even <-chan int, odd <-chan int, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("from the even channel", v)
		case v := <-odd:
			fmt.Println("from the even channel", v)
		case v := <-quit:
			fmt.Println("from the quit channel", v)
			return
		}

	}
}

func send(even chan<- int, odd chan<- int, quit chan<- int) {
	for i := 0; i < 100; i++ {
		if (i % 2) == 0 {
			even <- i
		}
		if (i % 2) == 1 {
			odd <- i
		}
	}

	quit <- 0
	close(even)
	close(odd)

}
