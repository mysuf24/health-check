package routes

import (
	"health-check/controller"
	"health-check/utils"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	utils.LogInfo("Registering routes")
	router.GET("/health", controller.HealthCheck)
}
