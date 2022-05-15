package main

import "fmt"

func main() {
	ch1 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	var v1 int
	for i := 0; i < 10; i++ {
		if i > 5 {
			break
		}
		select {
		case v1 = <-ch1:
			fmt.Println(v1)
		}
	}
}
