package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	var sum int = 0
	for _, num := range arr {
		sum += num
		fmt.Print(num)
	}
	println("\nSum of all elements in array is : ", sum)

	// range with map

	map1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}

	for k, v := range map1 {
		fmt.Println(k, v)
	}

	// range in strings
	for i, v := range "Hello" {
		fmt.Println(i, string(v))
	}
}
