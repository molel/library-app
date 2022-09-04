package handler

import (
	"github.com/gin-gonic/gin"
	"library-app/entities"
	"net/http"
)

func (h *Handler) SignUp(ctx *gin.Context) {
	var inputJSON entities.UserSignUp

	if err := ctx.BindJSON(&inputJSON); err != nil {
		ErrorResponse(ctx, err)
		return
	}

	userId, err := h.service.Auth.CreateUser(inputJSON)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"userId": userId})
}

func (h *Handler) SignIn(ctx *gin.Context) {
	var inputJSON entities.UserSignUp

	if err := ctx.BindJSON(&inputJSON); err != nil {
		ErrorResponse(ctx, err)
		return
	}

	token, err := h.service.Auth.GenerateToken(inputJSON.Username, inputJSON.Password)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"token": token})
}
