package users

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	db "github.com/JackleStyle0/micro_services_practice/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"log"
	"regexp"
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

	//conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can't connect db", err)
	}

	conPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to database")
	}

	query := db.New(conPool)
	user, err := query.GetUser(context.Background(), "yao3@example.com")
	if err != nil {
		fmt.Println(">>>> error", err)
	} else {
		fmt.Println(">>>> email", user.Email)
	}

	if err != nil {
		c.JSON(400, "invalid request1")
	}

	hash := md5.Sum([]byte(login.Password))
	hashPassword := hex.EncodeToString(hash[:])
	if user.Password != hashPassword {
		c.JSON(400, "invalid password")
	} else {
		c.JSON(200, user)
	}
}

func validateEmailPatten(email string) bool {
	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	isMatch := pattern.MatchString(email)
	return isMatch
}
