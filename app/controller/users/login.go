package users

import (
	"fmt"
	"regexp"

	"github.com/JackleStyle0/micro_services_practice/app/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var login models.UserLogin
	err := c.BindJSON(&login)
	if err != nil {
		fmt.Println(err)
	}

	if login.Email == "" || login.Password == "" {
		c.JSON(400, "Bad Request")
	}

	isValid := validateEmailPatten(login.Email)
	if isValid {
		c.JSON(200, "token")
	} else {
		c.JSON(400, "invalid request")
	}
}

func validateEmailPatten(email string) bool {
	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	isMatch := pattern.MatchString(email)
	return isMatch
}
