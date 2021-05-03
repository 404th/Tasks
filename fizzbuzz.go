package main

import "fmt"

func three(num int) bool {
	if num%3 == 0 {
		return true
	} else {
		return false
	}
}
func five(num int) bool {
	if num%5 == 0 {
		return true
	} else {
		return false
	}
}

func main() {

	var num int = 100

	for i := 0; i < num; i++ {
		if three(i) && five(i) {
			fmt.Println("FizzBuzz")
		} else if three(i) && !five(i) {
			fmt.Println("Fizz")
		} else if !three(i) && five(i) {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
