package main

import "fmt"

func main() {

	numbers := [...]int{100, 200, 300}
	size := len(numbers)
	fmt.Println("size of array =", size, numbers)
	fmt.Println(numbers)

}
