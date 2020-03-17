package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Demo() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("[DemoMiddleware] @ %v", time.Now())
		c.Next()
	}
}
