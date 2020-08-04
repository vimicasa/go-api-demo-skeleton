package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LogoutHandler logout function
func LogoutHandler(c *gin.Context) {
	c.AbortWithStatus(http.StatusOK)
}
