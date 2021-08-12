package main

import "fmt"

func main() {
	var baru = hello2("amalia")
	fmt.Println(baru)

	var name, age = hello3("amalia", 20) //kalau gamau age nya maka di parsing pakai _
	fmt.Println(name, age)

	listNumber(1, 2, 3)

	var closure1 = func(name string) { //closure atau anonymous
		fmt.Println("ini closure", "hello", name)
	}

	closure1("amal")
}

func hello1(name string) string {
	greetings := "hello " + name
	return greetings
}

func hello2(name string) string {
	return "hello " + name
}

func hello3(name string, age int) (string, int) {
	return "hallo " + name, age
}

func listNumber(number ...int) { //variadic function
	fmt.Println(number)
}
