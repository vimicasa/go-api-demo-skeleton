package login

import (
	r "bb-server/src/handler"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	User     string `json:"user"  binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	Message string `json:"Message"`
}

// Login generates the authentication token
// if the password was matched with the specified account.
func Login(c *gin.Context) {

	var json loginRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		r.BadRequestResponse(c, "binding_not_valid", "Binding is not valid", nil)
		return
	}

	if json.User != "vimicasa" || json.Password != "123456" {
		r.BadRequestResponse(c, "password_not_valid", "Password is not valid", nil)
		return
	}

	r.OKResponse(c, "Login Successfull", loginResponse{Message: "Welcome Test"})
}
