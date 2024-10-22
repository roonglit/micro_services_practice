package users

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"

	models "github.com/JackleStyle0/micro_services_practice/app/models/user"
	"github.com/gin-gonic/gin"
)

func LoginController(c *gin.Context) {
	mockEmail := "yao@odds.team"
	mockPassword := "382e0360e4eb7b70034fbaa69bec5786"

	var login models.Login
	err := c.BindJSON(&login)
	if err != nil {
		fmt.Println(err)
	}

	if login.Email == "" || login.Password == "" {
		c.JSON(400, "Bad Request")
	}

	isValid := validateEmailPatten(login.Email)
	if !isValid {
		c.JSON(400, "invalid request")
	}

	hash := md5.Sum([]byte(login.Password))
	hashPassword := hex.EncodeToString(hash[:])

	if login.Email == mockEmail && hashPassword == mockPassword {
		c.JSON(200, "OK")
	} else {
		c.JSON(400, "invalid request")
	}
}

func validateEmailPatten(email string) bool {
	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	isMatch := pattern.MatchString(email)
	return isMatch
}
