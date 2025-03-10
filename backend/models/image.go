package models

import (
	"time"
)

// Image represents the image model
type Image struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Filename    string    `json:"filename"`
	FilePath    string    `json:"filePath"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ImageDatabase is a simple in-memory database for images
// In a real application, you'd use a database like MySQL, PostgreSQL, or MongoDB
var ImageDatabase = make(map[string]*Image)