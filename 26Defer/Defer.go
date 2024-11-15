package main

import "fmt"

func add(value1, value2 float64) {
	result := value1 + value2
	fmt.Println("ผลลัพธ์ =", result)
}

func loop() {
	for a := 0; a < 10; a++ {
		fmt.Println("a =", a)
	}
}

func deferloop() {
	for b := 0; b < 10; b++ {
		defer fmt.Println("b =", b)
	}
}

func main() {
	fmt.Println("ยินดีต้อนรับ")
	defer fmt.Println("จบการทำงาน")
	add(20, 20)
	add(30, 20)

	loop()
	deferloop()
}
