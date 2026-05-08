package main

import "fmt"

// exercise1 creates 3 goroutines
// first two goroutines each send 10 numbers to a channel
// the third goroutine prints all 20 numbers
func exercise1(done chan struct{}) {
	out := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			out <- i + 1
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			out <- i + 11
		}
	}()

	go func() {
		defer close(out)
		for i := 0; i < 20; i++ {
			v := <-out
			fmt.Println(v)
		}
		done <- struct{}{}
	}()

}

func main() {
	done := make(chan struct{})
	exercise1(done)
	<-done
	fmt.Println("Done")
}
