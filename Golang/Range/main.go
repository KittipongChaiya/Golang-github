package main

import "fmt"

func main() {

	numbers2 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for index := 0; index < len(numbers2); index++ {
		fmt.Println(numbers2[index], index)
	}

	for index, value := range numbers2 {
		fmt.Println("Index = ", index, "Value = ", value)
	}

	language := map[string]string{"TH": "Thailand", "EN": "English", "JP": "Japan"}
	for key, value := range language {
		fmt.Println("Key =", key, "Value =", value)
	}
}
