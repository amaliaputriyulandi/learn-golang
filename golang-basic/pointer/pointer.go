package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	// var p1 int = 20
	// var p2 int = p1 //cuma nyimpen nilai p1 aja. jadi ketika p1 berubah tetap nilai p1 sebelumnya

	// //ini disebut variabel pointer untuk menyimpan alamat memori
	// var p3 *int = &p1 //ini menyimpan alamat memori. kalau p1 berubah maka nilainya p3 ikut ubah karna cuma nyimpem alamat memori
	// // & dibaca ampersand, p3 gabisa ubah nilai karena isinya itu alamat memori
	// //&p1 (tidak sama dengan nilai p1 = 20) (&p1 adalah alamat memori p1)

	// p3 = &p2 //kalau pointer hanya bisa mengubah alamat memoru

	// p1 = 30

	// fmt.Println(p1)
	// fmt.Println(p2)
	// fmt.Println(*p3) //cara liat value dari alamat memori

	var p5 = person{"agus", 24}
	var p6 = p5

	var p7 *person = &p5
	p5.name = "gugun"
	p7.name = "haikal"
	fmt.Println("p5", p5)
	fmt.Println("p6", p6)
	fmt.Println("p7", *p7)
}
