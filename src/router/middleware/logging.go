package middleware

import (
	u "bb-server/src/utils"
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/willf/pad"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Logging is a middleware function that logs the each request.
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		path := c.Request.URL.Path

		// // Read the Body content
		// var bodyBytes []byte
		// if c.Request.Body != nil {
		// 	bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		// }

		// // Restore the io.ReadCloser to its original state
		// c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// The basic informations.
		method := c.Request.Method
		ip := c.ClientIP()

		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw

		// Continue.
		c.Next()

		// Calculates the latency.
		end := time.Now().UTC()
		latency := end.Sub(start)

		statusCode := c.Writer.Status()
		u.InfoLogger.Printf("%s %s | %d | %s | %s ", pad.Right(method, 5, ""), path, statusCode, ip, latency)

		//Remove this. It is not necessary
		if statusCode >= 400 {
			//ok this is an request with error, let's make a record for it
			// now print body (or log in your preferred way)
			u.ErrorLogger.Println("Response body: " + blw.body.String())
		}
	}
}
