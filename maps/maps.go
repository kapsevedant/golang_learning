package main

import (
	"fmt"
)

func main() {
	map1 := make(map[string]int)
	fmt.Println(map1)

	map1["a"] = 1
	map1["b"] = 2
	map1["c"] = 3
	fmt.Println(map1["a"])
	fmt.Println(map1["d"]) // 0

	// delete method
	delete(map1, "b")
	fmt.Println(map1)

	// clear method make the map empty
	clear(map1)
	fmt.Println(map1)

	map2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(map2)

	v, ok := map2["d"]

	fmt.Println(v)  // 0
	fmt.Println(ok) // false

	v1, ok1 := map2["a"]
	fmt.Println(v1)  // 1
	fmt.Println(ok1) // true

	if ok {
		fmt.Println("key is present")
	} else {
		fmt.Println("key is not present")
	}

}
