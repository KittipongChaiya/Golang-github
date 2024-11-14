package main

import "fmt"

func main() {
	fmt.Println("เริ่มการทำงาน")

	fmt.Println("สวัสดีจาก Go!")

	name := "ทอม"
	age := 66
	score := 99.66
	ispass := true
	number1 := 10
	number2 := 3

	/*
		var name2 string
			ประกาศตัวแปร
			%s คือ placeholder สำหรับ string
		    %d คือ placeholder สำหรับตัวเลขจำนวนเต็ม
		    \n คือ การขึ้นบรรทัดใหม่
			%.2f คือ placeholder สำหรับจำนวนจริง float
	*/

	fmt.Printf("ฉันชื่อ %s และมีอายุ %d ปี คะแนน %.2f ผลสอบ %t\n", name, age, score, ispass)
	fmt.Printf("ฉันชื่อ %T และมีอายุ %T ปี คะแนน %T ผลสอบ %T\n", name, age, score, ispass)

	fmt.Println("ผลบวก =", number1+number2)
	fmt.Println("ผลลบ =", number1-number2)
	fmt.Println("ผลคูณ =", number1*number2)
	fmt.Println("ผลหาร =", number1/number2)
	fmt.Println("เศษ =", number1%number2)

	fmt.Println(number1, "กับ", number2, "ค่าเท่ากันหรือไม่ =", number1 == number2)
	fmt.Println(number1, "กับ", number2, "ค่าไม่เท่ากันหรือไม่ =", number1 != number2)
	fmt.Println(number1, "มากกว่า", number2, "หรือไม่ =", number1 > number2)
	fmt.Println(number1, "น้อยกว่า", number2, "หรือไม่ =", number1 < number2)
	fmt.Println(number1, "มากกว่าหรือเท่ากับ", number2, "หรือไม่ =", number1 >= number2)
	fmt.Println(number2, "น้อยกว่าหรือเท่ากับ", number2, "หรือไม่ =", number1 <= number2)

}
