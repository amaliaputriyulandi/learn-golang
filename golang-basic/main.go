package main //nama nya main dan wajib

import (
	"fmt" //lebih liatin data aja digunain
	"log"
)

func main() {
	//fmt.Println("hello world")
	const var1 = "Hello" //kaya biasa tapi bisa juga ga dipake
	var var2 string = "string" //kalau pake var harus dipake dan harus nentuin typenya
	var varNoValue string
	var var3 int = 223
	fmt.Println(var2)
	log.Println(var3)

	var4 := "var4 coba" //kaya var tapi gaperlu pake type, apa aja bisa. tapi harus dipake
	fmt.Println(var4)
	var4 = "uuu"
	fmt.Println(var4)
	fmt.Println(varNoValue)

	varNoValue = "jayaaa"
	fmt.Println(varNoValue)
}
