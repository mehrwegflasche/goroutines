package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		v1 := 2
		ch1 <- v1
		v2 := <-ch2
		fmt.Printf("%d,%d\n", v1, v2)
	}()
	v1 := 1
	var v2 int
	select {
	case ch2 <- v1:
	case v2 = <-ch1:
	}

	fmt.Printf("%d,%d\n", v1, v2)
}
