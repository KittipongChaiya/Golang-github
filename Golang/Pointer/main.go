package main

import "fmt"

func zerovalue(ivalue int) {
	ivalue = 0
}

func zeropointer(inpoiter *int) {
	*inpoiter = 0
}

func main() {
	i := 1
	fmt.Println(i)

	zerovalue(i)
	fmt.Println("i จากฟังชั่น zerovalue", i)

	zeropointer(&i)
	fmt.Println("i จากฟังชั่น zeropointer", i)

	zeropointer(&i)
	fmt.Println("i address จากฟังชั่น zeropointer", &i)
}
