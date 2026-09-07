package main

import "fmt"

const age = 23

func main() {
	fmt.Println("Age is ", age)

	const (
		port = 8080
		host = "localhost"
	)

	fmt.Println(port, host)
}
