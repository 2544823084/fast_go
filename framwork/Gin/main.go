package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := os.MkdirAll("logs", 0o755); err != nil {
		log.Fatalf("create logs dir: %v", err)
	}

	f, err := os.OpenFile("logs/gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		log.Fatalf("open log file: %v", err)
	}
	defer f.Close()

	// 同时写入控制台和 logs/gin.log
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
	gin.DefaultErrorWriter = io.MultiWriter(os.Stderr, f)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("run server: %v", err)
	}
}
