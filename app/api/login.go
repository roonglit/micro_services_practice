package api

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	_ "github.com/lib/pq"
)

func (server *Server) Login(c *gin.Context) {
	var login LoginReq
	err := c.ShouldBindJSON(&login)
	if err != nil {
		fmt.Println(err)
	}
	validate := validator.New()

	err = validate.Struct(login)
	if err != nil {
		// Validation failed, handle the error
		errors := err.(validator.ValidationErrors)
		fmt.Print("Validation error", errors)
		http.Error(c.Writer, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
		return
	}

	// if login.Email == "" || login.Password == "" {
	// 	c.JSON(400, "Bad Request")
	// }

	// isValid := validateEmailPatten(login.Email)
	// if !isValid {
	// 	c.JSON(400, "invalid request")
	// 	return
	// }

	//conn, err := sql.Open(dbDriver, dbSource)
	// if err != nil {
	// 	log.Fatal("can't connect db", err)
	// }

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

// func validateEmailPatten(email string) bool {
// 	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
// 	isMatch := pattern.MatchString(email)
// 	return isMatch

// }
