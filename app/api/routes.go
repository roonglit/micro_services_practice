package api

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.POST("/users/login", server.Login)
	return r
}
