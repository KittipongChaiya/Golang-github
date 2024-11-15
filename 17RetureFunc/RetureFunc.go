package main

import "fmt"

func main() {

	result := summation(500, 700, 400, 300, 200)
	fmt.Println("ยอดรวมทั้งหมด =", result)
}

func summation(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}
