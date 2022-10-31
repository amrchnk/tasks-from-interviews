package main

import "fmt"

// implement an algorithm that takes an array and moves zeros to the end,
// preserving the order of the rest of the elements

func main() {
	arr := []int{1, 2, 3, 0, 14, 5, 0, 6, 7, 0}
	newArr := ShiftZeros(arr)
	fmt.Println(newArr)
}

// позитивно

func ShiftZeros(arr []int) []int {
	var lastNonZeroIndex int

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[lastNonZeroIndex], arr[i] = arr[i], arr[lastNonZeroIndex]
			lastNonZeroIndex++
		}
	}

	return arr
}

// примитивно (какое О в таком случае?)

func ShiftZerosTwo(arr []int) []int {
	var tempZeros, tempArr []int
	for _, k := range arr {
		if k == 0 {
			tempZeros = append(tempZeros, k)
		} else {
			tempArr = append(tempArr, k)
		}
	}
	arr = append(tempArr, tempZeros...)
	return arr
}
