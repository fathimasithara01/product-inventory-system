package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Status: "success",
		Data:   data,
	})
}

func Message(c *gin.Context, message string) {
	c.JSON(200, Response{
		Status:  "success",
		Message: message,
	})
}

func Error(c *gin.Context, message string, code int) {
	c.JSON(code, Response{
		Status:  "error",
		Message: message,
	})
}
