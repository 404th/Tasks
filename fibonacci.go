package main

import "fmt"

func worker(rec <-chan int, send chan<- int) {
	for i := 0; i < 100; i++ {
		send <- fib(<-rec)
	}
}

func main() {

	c1 := make(chan int, 100)
	c2 := make(chan int, 100)

	go worker(c1, c2)

	for i := 0; i < 100; i++ {
		c1 <- i
	}
	close(c1)

	for i := 0; i < 100; i++ {
		fmt.Println(<-c2)
	}

}

// fibonacci sonini topish functioni
func fib(num int) int {

	if num <= 1 {
		return num
	}
	return fib(num-1) + fib(num-2)
}
