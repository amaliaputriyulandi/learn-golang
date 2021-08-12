package main

import "fmt"

func main() {
	arr1 := [3]int{1, 2, 3}
	string1 := "strings"

	//di golang untuk perulangan hanya bisa for saja
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	for i := 10; i >= 1; i-- {
		fmt.Println("Perulangan ke-", i)
	}

	for _, value := range string1 { //kalau mau value nya aja index nya jadi _ terus println gausah print index nya
		if string(value) == "s" {
			continue
			//fmt.Println("iya s")
		}
		//break
		fmt.Println("bukan s")
	}

	for {
		fmt.Println("perulangan")
		if 1 == 1 {
			break
		}
	}
}
