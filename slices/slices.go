package main

import (
	"fmt"
	"slices"
)

// -> SLice = Dynamic Array

func main() {

	// numbers := []int{1, 2, 3, 4, 5}
	// fmt.Printf("type of numbers is : %T\n", numbers)
	// fmt.Println(len(numbers))
	// numbers = append(numbers, 6, 7)
	// fmt.Println(numbers)
	// fmt.Println(len(numbers))

	// make method

	// arr := make([]int, 3, 5)
	// arr[0] = 10
	// arr = append(arr, 1, 2, 3, 4, 5, 6)
	// fmt.Println(arr)
	// fmt.Println(len(arr), cap(arr))

	// copy method
	// arr1 := make([]int, 0, 5)
	// arr1 = append(arr1, 1, 2)
	// arr2 := make([]int, len(arr1))

	// fmt.Println(arr1, arr2)

	// copy(arr2, arr1)
	// fmt.Println(arr1, arr2)

	// slice operator

	// var array = []int{1, 2, 3, 4, 5}
	// fmt.Println(array[1:4]) // [2 3 4]

	n1 := []int{1, 2}
	n2 := []int{1, 2}

	fmt.Println(slices.Equal(n1, n2))
}
