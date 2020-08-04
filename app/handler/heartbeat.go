package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HeartbeatHandler health check endpoint
func HeartbeatHandler(c *gin.Context) {
	c.AbortWithStatus(http.StatusOK)
}
