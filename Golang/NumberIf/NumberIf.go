package main

import "fmt"

func main() {
	var number3 int
	fmt.Print("ป้อนตัวเลข = ")
	fmt.Scanf("%d", &number3)

	if number3%2 == 0 {
		fmt.Println("เลขคู่")
	} else {
		fmt.Println("เลขคี่")
	}
}
