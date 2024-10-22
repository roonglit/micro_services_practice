package config

import (
	"github.com/JackleStyle0/micro_services_practice/app/controller/users"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) *gin.Engine {
	r.POST("/users/login", users.LoginController)
	return r
}
