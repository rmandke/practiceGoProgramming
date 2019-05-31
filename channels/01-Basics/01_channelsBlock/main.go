// This code does not work as channel locks access to the variable

package main

import "fmt"

func main() {

	c := make(chan int)

	c <- 42

	fmt.Println(<-c)

}
