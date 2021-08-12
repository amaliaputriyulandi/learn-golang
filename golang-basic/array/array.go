package main

import "fmt"

func main() {
	var arr1 = [5]int{1, 2, 3, 4, 5} //array harus ada length
	arr1[0] = 8

	arr2 := [...]int{1, 2, 3} //array

	slice1 := []int{1, 2, 3, 4, 5} //slice

	var arr3 = make([]int, 3) //3 itu size nya
	arr3[0] = 1

	fmt.Println(arr1)
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println("--------------------------------------------")
	fmt.Println(cap(slice1))    //cap itu kaya length hitung panjangnya
	slice1 = append(slice1, 10) //kaya push tapi hanya bisa di slice
	fmt.Println(len(slice1))    //len itu kaya length hitung panjangnya
	fmt.Println(arr3)

	fmt.Println("--------------------------------------------")
	slice2 := arr1[2:] // brrti dia ambil dari index 0 sampai sebelum index 2. kalau mau semua [:]. bisa juga [:2] brrti ambil dari index awal sampai sebelum index 2
	// [2:] brrti ambil dari index sampai selesai
	fmt.Println("slice2", slice2)
	fmt.Println("len slice2", len(slice2))
	fmt.Println("cap slice2", cap(slice2))
	fmt.Println("--------------------------------------------")
	slice2 = append(slice2, 2, 3)
	fmt.Println("slice2", slice2)
	fmt.Println("len slice2", len(slice2))
	fmt.Println("cap slice2", cap(slice2)) //capacity nya lenght buat menanmpung misal 3 ga ada sisa dia akan kali 2 jadi 6.
	fmt.Println("--------------------------------------------")
	slice2 = append(slice2, 7)
	fmt.Println("slice2", slice2)
	fmt.Println("len slice2", len(slice2))
	fmt.Println("cap slice2", cap(slice2)) //karna capacity nya masih muat jadi tetap 6

}
