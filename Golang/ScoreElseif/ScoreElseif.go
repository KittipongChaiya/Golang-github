package main

import "fmt"

func main() {
	var score3 int
	fmt.Print("กรุณาป้อนคะแนนนักเรียน = ")
	fmt.Scanf("%d", &score3)
	fmt.Println("สวัสดี =", score3)

	if score3 >= 90 {
		fmt.Println("เกรด A")
	} else if score3 >= 80 {
		fmt.Println("เกรด B")
	} else if score3 >= 70 {
		fmt.Println("เกรด C")
	} else {
		fmt.Println("ไม่ผ่าน")
	}
}
