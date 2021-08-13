package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUser(c *gin.Context) {
	var user User
	user.Username = "amaliaa"
	user.Password = "123456"

	c.JSON(http.StatusOK, user)
}

func PostUser(c *gin.Context) {
	var user User

	c.Bind(&user)

	fmt.Println(user.Username, user.Password)
	c.JSON(http.StatusOK, user)
}
