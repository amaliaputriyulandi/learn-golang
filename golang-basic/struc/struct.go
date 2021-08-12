package main

import "fmt"

type location struct {
	province string
	country  string
}

type contact struct {
	name    string
	email   string
	phone   string
	address location //bisa dibilang embed struct
}

type person struct {
	name    string
	age     int
	address struct {
		city string
	}
}

func main() {
	var c1 contact //c1 dengan type struct contact
	c1.name = "amalia"
	c1.email = "amalia@gmail.com"
	c1.phone = "082298006672"
	c1.address.province = "kepri"
	c1.address.country = "batam"

	var c2 = contact{ //kalau tidak pakai name: itu harus tulis berurutan kalau engga tidak berurut dan harus semua di isi
		name:    "amalias",
		phone:   "082298",
		address: location{country: "yogya", province: "DIY"},
	}

	fmt.Println(c1)
	fmt.Println(c2)

	//anonymous struct
	var c3 = struct {
		name string
		age  int
	}{} //{name: "putri", age: 8} or
	c3.name = "lulu"
	c3.age = 8

	var c4 = person{
		name: "ppp",
		age:  9,
		address: struct{ city string }{
			city: "batam",
		},
	}

	fmt.Println(c3)
	fmt.Println(c4)

	var listContact = []contact{
		{name: "lili", email: "amal@gmail.com", phone: "09988",
			address: location{province: "kepri", country: "batam"}},
		{name: "lili", email: "amal@gmail.com", phone: "09988",
			address: location{province: "kepri", country: "batam"}},
	}

	listContact = append(listContact, contact{name: "lili+", email: "amal@gmail.com", phone: "09988",
		address: location{province: "kepri", country: "batam"}})

	fmt.Println(listContact)

	c1.Hello()
}

//struct method
func (c contact) Hello() {
	fmt.Println("hello", c.name)
}
