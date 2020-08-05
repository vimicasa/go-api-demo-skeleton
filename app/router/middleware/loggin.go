package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/vimicasa/go-api-demo-skeleton/app"
)

// Logger is the logrus logger handler
func Logger() gin.HandlerFunc {

	return func(c *gin.Context) {
		// start time
		startTime := time.Now()
		// Processing request
		c.Next()
		// End time
		endTime := time.Now()
		// status
		status := c.Writer.Status()

		entry := app.LogAccess.WithFields(
			logrus.Fields{
				"status":   status,
				"method":   c.Request.Method,
				"path":     c.Request.RequestURI,
				"ip":       c.ClientIP(),
				"duration": endTime.Sub(startTime),
			})

		switch {
		case status > 499:
			entry.Error()
		case status > 399:
			entry.Warn()
		default:
			entry.Info()
		}

	}
}
