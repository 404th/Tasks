package main

import "fmt"

func main() {
	fmt.Println(fibo1(11))
	fmt.Println(fibo2(11))
}

// using recursive func
func fibo1(num int) int {
	if num <= 1 {
		return num
	}

	return fibo1(num-1) + fibo1(num-2)
}

// simple way
func fibo2(num int) int {
	x, y, c := 0, 1, 0
	for i := 0; i < num; i++ {
		c = y
		y += x
		x = c
	}

	return c
}
