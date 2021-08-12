package main

import "fmt"

func main() {
	var penjumlahan int = 10 + 22
	fmt.Println(penjumlahan)

	var pengurangan int = 3 - 1
	var perkalian int = 10 * 2

	num1 := 10
	num2 := 3
	num2 += 4

	fmt.Println("ini num2", num2)

	var result float64 = float64(num1) / float64(num2)

	fmt.Println(pengurangan)
	fmt.Println(perkalian)
	fmt.Println(result)

	var module int = 10 % 5
	fmt.Println(module)

	fmt.Println("OPERATOR LOGIKA")

	perbandingan1 := 10 != 9
	var perbandingan2 bool = 10 < 9
	var perbandingan3 bool = 10 <= 9

	fmt.Println(perbandingan1)
	fmt.Println(perbandingan2)
	fmt.Println(perbandingan3)

}
