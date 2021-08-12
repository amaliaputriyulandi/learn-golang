package hello

import "fmt"

func Greetings(name string) string { //nama function nya harus kapital di awal
	// fmt.Println("hello", name)
	return "hello " + name
	// fmt.Println(Penjumlahan(7, 7))
}

//namanya encapluased bisa akses function antar file
func Tambah(angka int) {
	fmt.Println(1 + angka)
}
