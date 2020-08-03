package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HeartbeatHandler(c *gin.Context) {
	c.AbortWithStatus(http.StatusOK)
}
