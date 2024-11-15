package main

import "os"

func main() {
	data1 := []byte("Hello666")
	err := os.WriteFile("C:/Users/WGT/Desktop/Golang/Golang-github/Golang/write/data.txt", data1, 0644)

	if err != nil {
		panic(err)
	}
	f, ferr := os.Create("C:/Users/WGT/Desktop/Golang/Golang-github/Golang/Write/emp.txt")
	if ferr != nil {
		panic(ferr)
	}

	defer f.Close()

	data2 := []byte("tom\n jim")
	os.WriteFile("emp.txt", data2, 0644)
}
