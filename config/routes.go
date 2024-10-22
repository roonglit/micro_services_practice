package config

import (
	api "github.com/JackleStyle0/micro_services_practice/app/api/users"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.POST("/users/login", api.Login)
	return r
}
