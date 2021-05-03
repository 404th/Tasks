package main

import "fmt"

func checkerArr(arr []int) []int {
	var newArr []int
	for i := 0; i < len(arr); i++ {
		exists := true
		if len(newArr) > 0 {
			for j := 0; j < len(newArr); j++ {
				if arr[i] == newArr[j] {
					exists = true
					break
				} else {
					exists = false
				}
			}
		} else {
			newArr = append(newArr, arr[0])
		}

		if !exists {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}

func main() {

	duplArr1 := []int{1, 23, 34, 2, 23, 4, 2, 1}
	duplArr2 := []int{23, 23, 34, 2, 42, 33, 4, 23, 4, 2, 1}
	duplArr3 := []int{0, 0, 12, 23, 0, 12, 23, 23, 34, 0, 2, 42, 0, 33, 4, 23, 4, 2, 1}

	fmt.Println(checkerArr(duplArr1))
	fmt.Println(checkerArr(duplArr2))
	fmt.Println(checkerArr(duplArr3))

}
