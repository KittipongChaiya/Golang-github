package main

import "fmt"

func main() {
	for count := 1; count <= 10; count++ {
		fmt.Println("Hello Tom", count)
	}
	for count := 10; count >= 1; count-- {
		fmt.Println("Hello Tom", count)
	}
}
