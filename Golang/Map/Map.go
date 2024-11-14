package main

import "fmt"

func main() {

	//country := map[string]string{}
	//country["TH"] = "ไทย"
	//country["JP"] = "ญี่ปุ่น"
	//fmt.Println(country["TH"])

	country := map[string]string{"TH": "ไทย", "JP": "ญี่ปุ่น"}
	fmt.Println(country["TH"])
	value, check := country["EN"]
	if check {
		fmt.Println(value)
	} else {
		fmt.Println("ไม่พบข้อมูล")
	}
}
