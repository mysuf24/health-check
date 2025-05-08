package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		duration := time.Since(start)

		log.Printf("[REQUEST] %s %s | %d | %v",
			ctx.Request.Method,
			ctx.Request.URL.Path,
			ctx.Writer.Status(),
			duration,
		)
	}
}
