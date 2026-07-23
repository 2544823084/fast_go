package main

import (
	"log"
	"os"

	"gobox/framwork/gin/middleware"
	"gobox/framwork/gin/router"
)

func main() {
	closer, err := middleware.Init()
	if err != nil {
		log.Fatalf("init logger: %v", err)
	}
	defer closer.Close()

	r := router.Setup()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("run server: %v", err)
	}
}
