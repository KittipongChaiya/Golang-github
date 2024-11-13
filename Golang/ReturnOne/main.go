package main

import "fmt"

func main() {

	total(50, 100)
	delivery := getdelivery()
	fmt.Println("ค่าจัดส่ง=", delivery)
}

func total(number6, number7 int) {
	fmt.Println("ยอดรวม", number6+number7)
}

func getdelivery() int {
	return 50
}
