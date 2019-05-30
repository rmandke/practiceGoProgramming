// The objective of this code is to create a race condition by having multiple go routines run in parallel
// and verify that the race condition exsist using the --race flag on go run statement

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var mu sync.Mutex

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var incrementor int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {

		go func() {

			atomic.AddInt64(&incrementor, 1)
			fmt.Println(atomic.LoadInt64(&incrementor))
			wg.Done()

		}()

	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", incrementor)

}
