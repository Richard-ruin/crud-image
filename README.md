# CRUD Image Upload Application

A simple CRUD (Create, Read, Update, Delete) application for uploading and managing images, built with Go (Golang) backend and Svelte frontend.

## Features

- Upload images with title and description
- View all uploaded images in a responsive grid
- Edit image metadata (title and description)
- Delete images
- Preview images before upload
- Responsive design with Tailwind CSS

## Tech Stack

### Backend
- **Go (Golang)**: Programming language
- **Gin**: Web framework for Go
- **In-memory database**: Simple storage for this demo (can be replaced with a real database)

### Frontend
- **Svelte**: Frontend framework
- **Tailwind CSS**: Utility-first CSS framework
- **Rollup**: Module bundler

## Project Structure

```
crud-image/
├── backend/
│   ├── main.go               # Main application entry point
│   ├── handlers/
│   │   └── image_handler.go  # HTTP handlers for image operations
│   ├── models/
│   │   └── image.go          # Image data model
│   ├── config/
│   │   └── database.go       # Database configuration (empty for now)
│   └── uploads/              # Directory to store uploaded images
│
└── frontend/
    ├── public/               # Public assets
    ├── src/
    │   ├── App.svelte        # Root Svelte component
    │   ├── main.js           # JavaScript entry point
    │   ├── app.css           # Tailwind CSS imports
    │   ├── components/       # Svelte components
    │   └── services/
    │       └── api.js        # API service functions
    ├── package.json          # NPM dependencies
    ├── rollup.config.js      # Rollup configuration
    ├── postcss.config.js     # PostCSS configuration
    └── tailwind.config.js    # Tailwind CSS configuration
```

## Getting Started

### Prerequisites

- Go (1.16+)
- Node.js (14+)
- npm

### Installation and Setup

1. Clone the repository (or download the files)

```bash
git clone https://github.com/yourusername/crud-image.git
cd crud-image
```

2. Set up the backend

```bash
cd backend

# Initialize Go module
go mod init github.com/yourusername/crud-image/backend

# Install dependencies
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors
go get github.com/google/uuid

# Run the server
go run main.go
```

3. Set up the frontend

```bash
cd ../frontend

# Install dependencies
npm install

# Run the development server
npm run dev
```

4. Open your browser and navigate to:
   - Frontend: http://localhost:5000
   - Backend API: http://localhost:8080/api

## API Endpoints

The backend provides the following RESTful API endpoints:

| Method | Endpoint        | Description                   |
|--------|-----------------|-------------------------------|
| GET    | /api/images     | Get all images                |
| GET    | /api/images/:id | Get a specific image by ID    |
| POST   | /api/images     | Upload a new image            |
| PUT    | /api/images/:id | Update an existing image      |
| DELETE | /api/images/:id | Delete an image               |

Static files are served from:
- /uploads/* - For accessing uploaded images

## Future Enhancements

- Add user authentication
- Implement a real database (PostgreSQL, MongoDB, etc.)
- Add image categories/tags
- Implement pagination for the image list
- Add search functionality
- Add image resizing/compression

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Tailwind CSS - For the utility-first CSS framework
- Gin - For the lightweight web framework for Go
- Svelte - For the reactive frontend framework

---

Created by Richard Joe - 2025
