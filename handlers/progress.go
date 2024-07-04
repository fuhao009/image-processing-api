package handlers

import (
	"image-processing-api/agent"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProgressHandler(c *gin.Context) {
	id := c.Query("id")

	agent.GlobalProgress.Mu.Lock()
	prog, exists := agent.GlobalProgress.Progress[id]
	agent.GlobalProgress.Mu.Unlock()

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
