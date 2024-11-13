package main

import "fmt"

func main() {

	var number5 int
	fmt.Print("ป้อนตัวเลข = ")
	fmt.Scanf("%d", &number5)
	switch number5 {
	case 1:
		fmt.Println("เปิดบัญชีใหม่")
	case 2:
		fmt.Println("ฝากเงิน - ถอนเงิน")
	default:
		fmt.Println("ข้อมูลไม่ถูกต้อง")
	}
}
