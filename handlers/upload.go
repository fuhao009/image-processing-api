package handlers

import (
	"image-processing-api/agent"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	inputDir := c.PostForm("input_dir")
	outputDir := c.PostForm("output_dir")
	id := c.PostForm("id")

	go agent.ProcessImages(inputDir, outputDir, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Processing started",
		"id":      id,
	})
}
