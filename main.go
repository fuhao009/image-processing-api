package main

import (
	"image-processing-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/upload", handlers.UploadHandler)
	r.GET("/progress", handlers.ProgressHandler)
	r.Run(":8080")
}
