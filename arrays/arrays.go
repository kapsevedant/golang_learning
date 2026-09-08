package main

import "fmt"

// if we have the fixed size then we can use the arrays, memory optimization, constant time access, but we cannot change the size of the array
// generally we use slices in go, slices are dynamic arrays, we can change the size of the slice, but it is slower than arrays

func main() {

	// var nums [5]int
	// nums[0] = 10
	// nums[1] = 20
	// fmt.Println(len(nums))
	// fmt.Println("Array is ", nums)

	// var values [4]bool

	// values[0] = true

	// fmt.Println("Bool array is ", values)

	var array = [5]int{10, 20, 30, 40, 50}

	for i := 0; i < len(array); i++ {
		fmt.Println("Value:", array[i])
	}
}
