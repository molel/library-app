package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Error struct {
	error
}

func (e Error) JSON() map[string]interface{} {
	return map[string]interface{}{"error": e.Error()}
}

func ErrorResponse(ctx *gin.Context, err error) {
	log.Printf("error occured: %s\n", err.Error())
	ctx.AbortWithStatusJSON(http.StatusBadRequest, Error{err}.JSON())
}
