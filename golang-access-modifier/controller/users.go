package controller

import (
	"fmt"
	"golang-access-modifier/model"
)

func CreateUser() {
	var user1 model.Users
	user1.Id = 1
	user1.Name = "Amalia"

	fmt.Println(user1.Id)
	fmt.Println(user1.Name)

}
