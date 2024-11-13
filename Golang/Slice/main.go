package main

import "fmt"

func main() {
	numbers := []int{100, 200, 300}
	numbers = append(numbers, 400)
	numbers = append(numbers, 500)
	fmt.Println(numbers[:])
	fmt.Println(numbers[:3])
}
