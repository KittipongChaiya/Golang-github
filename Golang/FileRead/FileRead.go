package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("C:/Users/WGT/Desktop/Golang/Golang-github/Golang/FileRead/Book1.csv")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	count := 1
	for scanner.Scan() {
		line := scanner.Text()
		item := strings.Split(line, ",")
		//fmt.Printf("line %d : %s\n", count, line)
		fmt.Println(item[2])

		count++
	}

}
