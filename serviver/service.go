package ccc

import (
	"fmt"
	"golang-access-modifier/hello"
)

func Punya(nama string) {
	coba := hello.Greetings(nama)

	fmt.Println(coba)
}
