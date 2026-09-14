package main

import "fmt"

func printNum(nums []int) {
	for _, item := range nums {
		fmt.Println("item value is ", item)
	}
}

func printStr(arr []string) {
	for _, item := range arr {
		fmt.Println("item value is ", item)
	}
}

func printArray[T any](arr []T) {
	for _, item := range arr {
		fmt.Println("item value is ", item)
	}
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []string{"a", "b", "c", "d", "e"}
	printNum(arr1)
	printStr(arr2)
	printArray(arr1)
	printArray(arr2)
}
