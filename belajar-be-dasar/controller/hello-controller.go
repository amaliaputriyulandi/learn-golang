package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Username string
	Password string
}

func Hello(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(rw, "Bad Request", 400)
		return
	}

	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	//localhost:9000/hello?name=amalia&age=10
	fmt.Fprintln(rw, "Hello, my name is ", name, "My age", age) //ini gabisa dijadi in json
}

//klik ada ResponseWriter untuk liat function apa aja di server.go
func GetUser(rw http.ResponseWriter, r *http.Request) {
	var user User
	user.Username = "amalia"
	user.Password = "amalia34"

	//dari struct ke Marshal
	dataByte, err := json.Marshal(user) //untuk ubah ke []byte
	if err != nil {
		fmt.Println("Cannot marshal data", err)
	}

	//var str string = "amaliaaputri"
	rw.Write(dataByte)
	// rw.Write([]byte(str))
}

func PostUser(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(rw, "Bad Request", 400)
		return
	}

	var user User

	//kalau ngubah dari byte ke struct itu unMarshal
	userByte, _ := ioutil.ReadAll(r.Body) //harus ngebaca dlu semua di body nanti tarus di variabel userByte
	err := json.Unmarshal(userByte, &user)
	if err != nil {
		fmt.Println("Cannot Unmarshall data", err)
	}

	rw.Write(userByte)

	// fmt.Println(user.Username, user.Password)

}
