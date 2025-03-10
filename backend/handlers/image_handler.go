package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yourusername/crud-image/backend/models"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// GetImages returns all images
func GetImages(c *gin.Context) {
	var images []*models.Image
	for _, img := range models.ImageDatabase {
		images = append(images, img)
	}
	
	// Always return an array, even if empty
	if images == nil {
		images = []*models.Image{}
	}
	
	c.JSON(http.StatusOK, images)
}

// GetImage returns a specific image by ID
func GetImage(c *gin.Context) {
	id := c.Param("id")
	if img, exists := models.ImageDatabase[id]; exists {
		c.JSON(http.StatusOK, img)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
}

// UploadImage handles image upload
func UploadImage(c *gin.Context) {
	// Parse form data
	title := c.PostForm("title")
	description := c.PostForm("description")

	// Get the file from form data
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Generate a unique filename
	extension := filepath.Ext(file.Filename)
	newFilename := uuid.New().String() + extension
	filePath := filepath.Join("uploads", newFilename)

	// Save the file
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Create image record
	now := time.Now()
	id := uuid.New().String()
	newImage := &models.Image{
		ID:          id,
		Title:       title,
		Filename:    newFilename,
		FilePath:    "/uploads/" + newFilename, // The URL path to access the image
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Save to database
	models.ImageDatabase[id] = newImage

	c.JSON(http.StatusCreated, newImage)
}

// UpdateImage updates an existing image's metadata
func UpdateImage(c *gin.Context) {
	id := c.Param("id")
	img, exists := models.ImageDatabase[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	// Update fields
	if title := c.PostForm("title"); title != "" {
		img.Title = title
	}
	if description := c.PostForm("description"); description != "" {
		img.Description = description
	}

	// Update timestamp
	img.UpdatedAt = time.Now()

	c.JSON(http.StatusOK, img)
}

// DeleteImage removes an image
func DeleteImage(c *gin.Context) {
	id := c.Param("id")
	img, exists := models.ImageDatabase[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	// Remove file from filesystem
	// Note: In a production app, you might want to handle errors more gracefully
	filePath := filepath.Join(".", img.FilePath)
	if err := os.Remove(filePath); err != nil {
		fmt.Printf("Warning: Could not delete file %s: %v\n", filePath, err)
	}

	// Remove from database
	delete(models.ImageDatabase, id)

	c.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
}