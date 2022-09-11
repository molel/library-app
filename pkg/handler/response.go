package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

// TODO change Error answer structure to return status and string message

type Error struct {
	error
}

func (e Error) JSON() map[string]interface{} {
	return map[string]interface{}{"error": e.Error()}
}

func ErrorResponse(ctx *gin.Context, status int, err error) {
	log.Printf("error occured: %s\n", err.Error())
	ctx.AbortWithStatusJSON(status, Error{err}.JSON())
}
