package users

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	db "github.com/JackleStyle0/micro_services_practice/db/sqlc"
	"log"
	"regexp"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:root@localhost:5432/auth_development?sslmode=disable"
)

func LoginController(c *gin.Context) {
	var login db.User
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

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can't connect db", err)
	}
	queries := db.New(conn)
	res, err := queries.GetUser(context.Background(), login.Email)

	if err != nil {
		c.JSON(400, "invalid request1")
	}

	hash := md5.Sum([]byte(login.Password))
	hashPassword := hex.EncodeToString(hash[:])
	if res.Password != hashPassword {
		c.JSON(400, "invalid password")
	} else {
		c.JSON(200, res)
	}
}

func validateEmailPatten(email string) bool {
	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	isMatch := pattern.MatchString(email)
	return isMatch
}
