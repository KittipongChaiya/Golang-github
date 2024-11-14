package main

import "fmt"

func main() {
	var score2 float32
	fmt.Print("กรุณาป้อนคะแนนนักเรียน = ")
	fmt.Scanf("%f", &score2)
	fmt.Println("คะแนนสอบ + (จิตพิสัย)=", score2+10)

}
