package main

import (
	"fmt"
	"sync"
)

func main() {

	var wt sync.WaitGroup
	wt.Add(1)
	go fibo1(10, &wt)
	// go fmt.Println(fibo2(101, &wt))
	wt.Wait()
}

// using recursive func
func fibo1(num int, wt *sync.WaitGroup) int {
	if num <= 1 {
		fmt.Println(num)
		return num
	}

	fmt.Println(num)
	wt.Done()
	return fibo1(num-1, wt) + fibo1(num-2, wt)
}

// simple way
// func fibo2(num int, wt *sync.WaitGroup) int {
// 	defer wt.Done()
// 	wt.Wait()
// 	x, y, c := 0, 1, 0
// 	for i := 0; i < num; i++ {
// 		c = y
// 		y += x
// 		x = c
// 	}

// 	return c
// }
