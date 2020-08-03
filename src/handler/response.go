package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status      int         `json:"status"`
	Code        string      `json:"code,omitempty"`
	Description string      `json:"description,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

func CustomResponse(c *gin.Context, status int, code string, description string, data interface{}) {

	c.JSON(status, Response{
		Status:      status,
		Code:        code,
		Description: description,
		Data:        data,
	})
}

func NotFoundResponse(c *gin.Context) {

	c.JSON(http.StatusNotFound, Response{
		Status: http.StatusNotFound,
		Code:   "not_found",
	})
}

func BadRequestResponse(c *gin.Context, code string, description string, data interface{}) {

	c.JSON(http.StatusBadRequest, Response{
		Status:      http.StatusBadRequest,
		Code:        code,
		Description: description,
		Data:        data,
	})
}

func OKResponse(c *gin.Context, description string, data interface{}) {

	c.JSON(http.StatusOK, Response{
		Status:      http.StatusOK,
		Description: description,
		Data:        data,
	})
}
