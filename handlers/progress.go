package handlers

import (
	"github.com/gin-gonic/gin"
	"image-processing-api/agent"
	"net/http"
)

func ProgressHandler(c *gin.Context) {
	id := c.Query("id")

	agent.Progress.mu.Lock()
	prog, exists := agent.Progress.progress[id]
	agent.Progress.mu.Unlock()

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
