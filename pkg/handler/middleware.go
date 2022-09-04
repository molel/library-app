package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) UserIdentification(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		ErrorResponse(ctx, http.StatusUnauthorized, errors.New("empty authorization header"))
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		ErrorResponse(ctx, http.StatusUnauthorized, errors.New("invalid authorization header"))
		return
	}
	if len(headerParts[1]) == 0 {
		ErrorResponse(ctx, http.StatusUnauthorized, errors.New("empty token"))
		return
	}
	userId, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		ErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}
	ctx.Set("userId", userId)
}
