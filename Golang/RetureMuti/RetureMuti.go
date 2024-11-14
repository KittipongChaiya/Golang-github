package main

import "fmt"

func main() {

	mycart, status := gettotalcart(500, 700)
	fmt.Println("ยอดรวมทั้งหมด =", mycart, status)
}

func gettotalcart(number8, number9 int) (int, string) {
	total := number8 + number9
	status := ""
	if total%2 == 0 {
		status = "เลขคู่"
	} else {
		status = "เลขคี่"
	}
	return total, status
}
