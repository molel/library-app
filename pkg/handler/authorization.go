package handler

import (
	"github.com/gin-gonic/gin"
	"library-app/entities"
	"net/http"
)

func (h *Handler) SignUp(ctx *gin.Context) {
	var inputJSON entities.UserCreate
	if err := ctx.BindJSON(&inputJSON); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	userId, err := h.service.Authorization.CreateUser(inputJSON)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"id": userId})
}

func (h *Handler) SignIn(ctx *gin.Context) {
	var inputJSON entities.UserCreate
	if err := ctx.BindJSON(&inputJSON); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	token, err := h.service.Authorization.GenerateToken(inputJSON.Username, inputJSON.Password)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"token": token})
}
