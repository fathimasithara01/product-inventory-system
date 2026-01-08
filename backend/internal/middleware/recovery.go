package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Panic recovered:", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  "error",
					"message": "internal server error",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
