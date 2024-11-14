package main

import "fmt"

func main() {

	var number4 int
	fmt.Print("ป้อนตัวเลข = ")
	fmt.Scanf("%d", &number4)
	if number4 == 1 {
		fmt.Println("เปิดบัญชีใหม่")
	} else if number4 == 2 {
		fmt.Println("ฝากเงิน - ถอนเงิน")
	} else {
		fmt.Println("ข้อมูลไม่ถูกต้อง")
	}
}
