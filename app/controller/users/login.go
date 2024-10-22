package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	email := c.DefaultQuery("email", "")
	password := c.DefaultQuery("password", "")
	fmt.Println(email, password)
}
