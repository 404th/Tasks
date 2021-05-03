package main

import "fmt"

func center(arr []int) int {
	if len(arr)%2 == 0 {
		return len(arr) / 2
	} else {
		return (len(arr) - 1) / 2
	}
}

//
func extractNumber(num int) []int {

	var digits []int

	for num != 0 {
		remainder := num % 10
		num /= 10
		digits = append(digits, remainder)
	}

	return digits
}

func checker(arr []int) []bool {
	var isTrue []bool
	for i := 0; i < center(arr); i++ {
		if arr[i] == arr[len(arr)-(i+1)] {
			isTrue = append(isTrue, true)
		} else {
			isTrue = append(isTrue, false)
		}
	}

	return isTrue
}

func isPalendrom(x []bool, originArr []int) bool {
	var isP bool = true

	for i := 0; i < center(originArr); i++ {
		if !x[i] {
			isP = false
			break
		}
	}
	return isP
}

func main() {
	////
	digitsArr := extractNumber(12321)
	digitsIsTrue := checker(digitsArr)

	if isPalendrom(digitsIsTrue, digitsArr) {
		fmt.Println("Yes, it is Palendrome!")
	} else {
		fmt.Println("No, it is not Palendrome!")
	}
}
