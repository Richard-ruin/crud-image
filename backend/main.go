package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yourusername/crud-image/backend/handlers"
	"log"
	"os"
)

func main() {
	// Create uploads directory if it doesn't exist
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		err := os.Mkdir("uploads", 0755)
		if err != nil {
			log.Fatalf("Failed to create uploads directory: %v", err)
		}
	}

	// Initialize Gin router
	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Serve static files from uploads directory
	r.Static("/uploads", "./uploads")

	// Define API routes
	api := r.Group("/api")
	{
		api.GET("/images", handlers.GetImages)
		api.GET("/images/:id", handlers.GetImage)
		api.POST("/images", handlers.UploadImage)
		api.PUT("/images/:id", handlers.UpdateImage)
		api.DELETE("/images/:id", handlers.DeleteImage)
	}

	// Start server
	log.Println("Server starting on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}