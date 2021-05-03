package main

import "fmt"

func main() {
	var num1 int = 12345
	var num2 int = 345
	var num3 int = 45
	var num4 int = 12323123
	// test arr
	// var arr = []int{12, 23, 2, 4, 6, 9, 3}

	fmt.Println(checkerSumm(extractNumber(num1)))
	fmt.Println(checkerSumm(extractNumber(num2)))
	fmt.Println(checkerSumm(extractNumber(num3)))
	fmt.Println(checkerSumm(extractNumber(num4)))
}

func checkerSumm(arr []int) []int {

	var oddSumm int = 0
	var evenSumm int = 0

	for i := 1; i < len(arr); i++ {
		if i%2 == 0 {
			evenSumm += arr[i]
		} else {
			oddSumm += arr[i]
		}
	}

	send := []int{}

	return append(send, oddSumm, evenSumm+arr[0])
}

func extractNumber(num int) []int {

	var digits []int

	for num != 0 {
		remainder := num % 10
		num /= 10
		digits = append(digits, remainder)
	}

	for x, y := 0, len(digits)-1; y > x; x, y = x+1, y-1 {
		digits[x], digits[y] = digits[y], digits[x]
	}

	return digits
}
