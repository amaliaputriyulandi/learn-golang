package main

import "fmt"

func main() {
	map1 := map[string]string{
		"name": "agus",
		"age":  "24",
	}

	map2 := map[string]interface{}{
		"name": "amal",
		"age":  20,
	}

	map3 := map[string]int{
		"name": 22,
		"age":  20,
	}

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3["name"])
}
