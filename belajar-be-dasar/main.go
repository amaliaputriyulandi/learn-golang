package main

import (
	"belajar-be-dasar/controller"
	"fmt"
	"net/http"
)

// func hello(rw http.ResponseWriter, r *http.Request) {
// 	if r.Method != "GET" {
// 		http.Error(rw, "Bad Request", 400)
// 		return
// 	}
// 	fmt.Fprintln(rw, "Hello, my name is amalia")
// }

func main() {
	//route buat api nanti http.HandleFunc
	http.HandleFunc("/coba", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Hello, This is my first api in golang")
	})
	// app.GET("", )

	http.HandleFunc("/hello", controller.Hello)
	http.HandleFunc("/getuser", controller.GetUser)
	http.HandleFunc("/postuser", controller.PostUser)
	//salah satu struct method untuk nge passing data seperti hello di atas

	fmt.Println("Server is running on port 9000")
	http.ListenAndServe("localhost:9000", nil)
}
