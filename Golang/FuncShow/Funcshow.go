package main

import "fmt"

func main() {
	showMessage("tom")
	showMessage("wow")
	showMessage("aoo")
}

func showMessage(name string) {
	fmt.Println("Hello", name)
}
