package main

import "fmt"

func catch() {
	var r = recover()
	if r != nil {
		fmt.Println("error", r)
	} else {
		fmt.Println("Aplikasi ")
	}
}

func main() {
	defer catch()

	panic("error panic")
}
