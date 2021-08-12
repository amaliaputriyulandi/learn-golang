package main

import (
	"fmt"
	"strconv"
)

func main() {
	var string1 string = "10"
	var string2 string = "5"
	//fmt.Println(string1)
	numStr1, err := strconv.ParseInt(string1, 10, 32) //harus ada 2 variabel, dan ditetapin jika error harus print apaaaa gtu
	if err != nil {
		fmt.Println("error parsing", err)
	}

	numStr2, _ := strconv.ParseInt(string2, 10, 32)
	fmt.Println(numStr1)
	fmt.Println(numStr2)
	fmt.Println(numStr2 + numStr1)
}
