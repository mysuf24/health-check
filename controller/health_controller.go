package controller

import (
	"health-check/config"
	"health-check/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	utils.LogInfo("Call endpoint HealthCheck")

	err := config.PingDB()
	if err != nil {
		utils.LogError("Database tidak terhubung")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Database tidak terhubung",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "API aman dan database terhubung",
	})
}
