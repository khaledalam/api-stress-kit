package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/gin-gonic/gin"

	"api-stress-kit/routes"
)

func main() {

	go func() { http.ListenAndServe("0.0.0.0:6060", nil) }() // pprof

	r := gin.Default()

	// Basic ping endpoint with latency measurement
	r.GET("/ping", func(c *gin.Context) {
		start := time.Now()
		c.JSON(http.StatusOK, gin.H{
			"message":    "pong",
			"latency_ms": time.Since(start).Milliseconds(),
		})
	})

	// Register additional routes
	routes.RegisterRoutes(r)

	// Start server
	r.Run(":8080")
}
