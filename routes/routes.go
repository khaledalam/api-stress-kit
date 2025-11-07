package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api-stress-kit/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/users", handlers.GetUsers)
}
