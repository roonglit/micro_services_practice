package main

import (
	"github.com/JackleStyle0/micro_services_practice/config"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r = config.SetupRoutes(r)
	r.Run("localhost:3000")
}
