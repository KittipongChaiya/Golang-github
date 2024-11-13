package main

import (
	"fmt" //นำ package input/output เข้าทำงาน
	calculator "golang/Calculator"
)

func main() {
	fmt.Println("เริ่มการทำงาน")

	fmt.Println(calculator.Add(50, 100))
	fmt.Println(calculator.Subtract(100, 50))

}
