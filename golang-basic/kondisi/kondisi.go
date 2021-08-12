package main

import "fmt"

func main() {
	var1 := 1
	var2 := 2

	if var1 == var2 {
		fmt.Println("sama")
	} else {
		fmt.Println("beda")
	}

	switch var1 { //buat perbandingan var1 sama tidak dengan 1 atau 2
	case 1:
		fmt.Println("benar")
	case 2:
		fmt.Println("salah")
	}

	switch { //buat kondisi
	case (var1 == 1):
		fmt.Println("salah")
	case (var2 != var1):
		fmt.Println("salah")
	case (var2 == var1):
		fmt.Println("benar")
	}
}
