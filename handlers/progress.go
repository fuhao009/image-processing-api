package handlers

import (
	"github.com/gin-gonic/gin"
	"image-processing-api/agent"
	"net/http"
)

func ProgressHandler(c *gin.Context) {
	id := c.Query("id")

	agent.GlobalProgress.mu.Lock()
	defer agent.GlobalProgress.mu.Unlock()
	prog, exists := agent.GlobalProgress.Progress[id]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ID not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       id,
		"progress": prog,
	})
}
