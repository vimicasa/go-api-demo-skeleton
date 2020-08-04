package handler

import (
	"go-api-demo-skeleton/app"
	"go-api-demo-skeleton/app/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// TODO: In memory user
var dummyUser = user{
	ID:       "155",
	Username: "test",
	Password: "test",
}

type userRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginHandler handler user/password
func LoginHandler(c *gin.Context) {
	var u userRequest
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	//compare the user from the request, with the one we defined:
	if dummyUser.Username != u.Username || dummyUser.Password != u.Password {
		c.AbortWithStatusJSON(http.StatusUnauthorized, app.Response{
			Status:      http.StatusUnauthorized,
			Description: "Invalid credentials"})
		return
	}

	token, err := auth.CreateToken(dummyUser.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, app.Response{
			Status:      http.StatusInternalServerError,
			Description: "Internal error"})
		return
	}

	c.JSON(http.StatusOK, token)
}
