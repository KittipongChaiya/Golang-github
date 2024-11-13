package main

import (
	"fmt" //นำ package input/output เข้าทำงาน
)

type product struct {
	name     string
	price    float64
	category string
	discount int
}

func main() {
	fmt.Println("เริ่มการทำงาน")

	product1 := product{name: "ปากกา", price: 50.55, category: "เครื่องเขียน", discount: 10}
	product2 := product{name: "เมาส์", price: 60.55, category: "อุปกรณ์คอม", discount: 20}
	product3 := product{name: "หน้าจอ", price: 70.55, category: "คอม", discount: 30}
	fmt.Println(product1)
	fmt.Println(product2)
	fmt.Println(product3)
}
