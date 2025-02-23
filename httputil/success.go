package httputil

import "github.com/gin-gonic/gin"

func NewSuccess(ctx *gin.Context, status int, message string) {
	er := HTTPSuccess{
		Code:    status,
		Message: message,
	}
	ctx.JSON(status, er)
}

type HTTPSuccess struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"status ok"`
}
