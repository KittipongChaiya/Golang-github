package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	go f("Hello Tom")
	go f("message2")
	time.Sleep(5 * time.Second)
}
