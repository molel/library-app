package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type ResponseStruct struct {
	Status  string
	Message string
}

func (ers ResponseStruct) JSON() map[string]interface{} {
	return map[string]interface{}{"status": ers.Status, "message": ers.Message}
}

func ErrorResponse(ctx *gin.Context, status int, err error) {
	log.Printf("error occured: %s\n", err.Error())
	ctx.AbortWithStatusJSON(status, ResponseStruct{Status: "error", Message: err.Error()}.JSON())
}
