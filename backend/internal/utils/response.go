package utils

import "github.com/gin-gonic/gin"

// Response defines a standard JSON response
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Success sends a JSON success response with data
func Success(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Status: "success",
		Data:   data,
	})
}

// Message sends a JSON success response with only a message
func Message(c *gin.Context, message string) {
	c.JSON(200, Response{
		Status:  "success",
		Message: message,
	})
}

// Error sends a JSON error response with a custom status code
func Error(c *gin.Context, message string, code int) {
	c.JSON(code, Response{
		Status:  "error",
		Message: message,
	})
}
