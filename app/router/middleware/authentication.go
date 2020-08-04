package middleware

import (
	"go-api-demo-skeleton/app"
	"go-api-demo-skeleton/app/auth"
	"go-api-demo-skeleton/app/model"
	"go-api-demo-skeleton/app/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthenticationRequired With list of roles
func AuthenticationRequired(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenStr := extractToken(c.Request)
		if tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, app.Response{
				Status:      http.StatusUnauthorized,
				Description: "Empty Token"})
			return
		}

		claims, err := auth.GetValidToken(tokenStr)
		userID, ok := claims["user_id"].(string)
		if err != nil || !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, app.Response{
				Status:      http.StatusUnauthorized,
				Description: "Invalid Token"})
			return
		}

		// Check if the user exists
		user, found := model.GetUser(userID)
		if !found {
			c.AbortWithStatusJSON(http.StatusUnauthorized, app.Response{
				Status:      http.StatusUnauthorized,
				Description: "UserToken Invalid "})
			return
		}

		if len(roles) != 0 {
			// Check roles
			_, found := util.Find(roles, user.Role)
			if !found {
				c.AbortWithStatusJSON(http.StatusForbidden, app.Response{
					Status:      http.StatusForbidden,
					Description: "Forbidden access"})
				return
			}
		}

		c.Next()

	}
}

//get the token from the Authorization header
func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
