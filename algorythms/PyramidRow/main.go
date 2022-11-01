package main

import "fmt"

//		1
// 	   3 5
//    7 9 11
//  13 15 17 19
// 21 23 25 27 29

// calculate the sum of a row of a pyramid of odd numbers (start from 1)
func main() {
	pyramid := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29}
	sum := PyramidRowSum(2, pyramid)
	fmt.Println("sum = ", sum)
}

// сложность = k

func PyramidRowSum(row int, arr []int) int {
	i := row * (row - 1) / 2
	var k, sum int
	for k != row {
		sum += arr[i+k]
		k++
	}
	return sum
}
