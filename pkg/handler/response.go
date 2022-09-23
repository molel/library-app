package handler

import (
	"github.com/gin-gonic/gin"
	"library-app/entities"
	"log"
)

func ErrorResponse(ctx *gin.Context, status int, err error) {
	log.Printf("error occured: %s\n", err.Error())
	ctx.AbortWithStatusJSON(status, entities.ResponseStruct{Status: "error", Message: err.Error()})
}
