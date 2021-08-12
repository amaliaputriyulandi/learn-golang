package main

import (
	"fmt"
	"golang-access-modifier/controller"
	"golang-access-modifier/hello"
)

func main() {
	cb := hello.Greetings("amalia")
	// fmt.Println("hello")
	// hello.Tambah(2)
	fmt.Println(cb)

	controller.CreateUser()
}
