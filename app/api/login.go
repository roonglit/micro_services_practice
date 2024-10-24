package api

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"regexp"

	db "github.com/JackleStyle0/micro_services_practice/db/sqlc"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func (server *Server) Login(c *gin.Context) {
	var login db.User
	err := c.ShouldBindJSON(&login)
	if err != nil {
		fmt.Println(err)
	}

	if login.Email == "" || login.Password == "" {
		c.JSON(400, "Bad Request")
	}

	isValid := validateEmailPatten(login.Email)
	if !isValid {
		c.JSON(400, "invalid request")
		return
	}

	//conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can't connect db", err)
	}

	user, err := server.store.GetUser(c, login.Email)
	if err != nil {
		fmt.Println(">>>> error", err)
	} else {
		fmt.Println(">>>> email", user.Email)
	}

	if err != nil {
		c.JSON(400, "invalid request1")
		return
	}

	hash := md5.Sum([]byte(login.Password))
	hashPassword := hex.EncodeToString(hash[:])
	if user.Password != hashPassword {
		c.JSON(400, "invalid password")
		return
	} else {
		c.JSON(200, user)
	}
}

func validateEmailPatten(email string) bool {
	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	isMatch := pattern.MatchString(email)
	return isMatch
}
