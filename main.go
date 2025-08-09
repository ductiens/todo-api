package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ductiens/todo-api/routes"
)

func main() {
	r := gin.Default()
	routes.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback khi chạy local
	}

	log.Println("Server is running on port", port)
	r.Run(":" + port)
}
