package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	ID           int
	EmployeeName string
	Phone        string
	Email        string
}

func main() {

	e := employee{}
	err := json.Unmarshal([]byte(`{"ID":1001,"EmployeeName":"Tom Banachek","Phone":"123456789","Email":"Tom@gmail.com"}`), &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e)
	fmt.Println(e.EmployeeName)
}
