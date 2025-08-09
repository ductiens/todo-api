package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/ductiens/todo-api/routes"
	"github.com/gin-contrib/cors"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	// Setup routes
	routes.RegisterRoutes(r)
	
	log.Println("Server is running on http://localhost:8080")
	r.Run(":8080")
}