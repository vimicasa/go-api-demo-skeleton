package handler

import (
	"go-api-demo-skeleton/app"
	"go-api-demo-skeleton/app/auth"
	"go-api-demo-skeleton/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginHandler handler user/password
func LoginHandler(c *gin.Context) {
	var u UserRequest
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	//compare the user from the request, with the one we defined:
	user, found := model.ValidUsernameAndPassword(u.Username, u.Password)
	if !found {
		c.AbortWithStatusJSON(http.StatusUnauthorized, app.Response{
			Status:      http.StatusUnauthorized,
			Description: "Invalid credentials"})
		return
	}

	token, err := auth.CreateToken(user.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, app.Response{
			Status:      http.StatusInternalServerError,
			Description: "Internal error"})
		return
	}

	c.JSON(http.StatusOK, token)
}
