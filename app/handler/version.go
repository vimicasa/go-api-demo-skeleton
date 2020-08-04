package handler

import (
	"go-api-demo-skeleton/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

// VersionHandler to retrieve version
func VersionHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"source":  "Test",
		"version": app.GetVersion(),
	})
}
