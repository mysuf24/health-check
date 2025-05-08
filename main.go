package main

import (
	"health-check/config"
	"health-check/middleware"
	"health-check/routes"
	"health-check/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LogInfo("memulai aplikasi...")

	config.ConnectDB()
	router := gin.Default()
	router.Use(middleware.RequestLogger())

	routes.RegisterRoutes(router)
	utils.LogInfo("Server berjalan di port 8080")
	router.Run(":8080")
}