package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	ID           int
	EmployeeName string
	Phone        string
	Email        string
}

func main() {
	data, _ := json.Marshal(&employee{1001, "Tom Banachek", "123456789", "tom@gmail.com"})
	fmt.Println(string(data))
}
