package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string) float64 {
	fmt.Printf("%v", promt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message, _ := fmt.Scanf("%v must number only", promt)
		panic(message)
	}
	return value
}

func add(value1, value2 float64) float64 {
	return value1 + value2
}

func add2(value1, value2 float64) float64 {
	return value1 - value2
}

func add3(value1, value2 float64) float64 {
	return value1 * value2
}

func add4(value1, value2 float64) float64 {
	return value1 / value2
}

func getOperator() string {
	fmt.Print("คำนวณได้แค่ ( + - * / ) : ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func main() {
	value1 := getInput("value1 =")
	value2 := getInput("value2 =")

	var result float64

	switch operator := getOperator(); operator {
	case "+":
		result = add(value1, value2)
	case "-":
		result = add2(value1, value2)
	case "*":
		result = add3(value1, value2)
	case "/":
		result = add4(value1, value2)
	default:
		panic("ใส่อะไรสักอย่างผิด")
	}

	fmt.Println(result)
}
