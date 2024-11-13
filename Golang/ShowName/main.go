package main

import (
	"fmt" //นำ package input/output เข้าทำงาน
)

func main() {
	fmt.Println("เริ่มการทำงาน")

	name2 := ""
	fmt.Print("กรุณาป้อนชื่อนักเรียน = ")
	fmt.Scanf("%s", &name2)
	fmt.Println("สวัสดี =", name2)
}
