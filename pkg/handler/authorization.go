package handler

import (
	"github.com/gin-gonic/gin"
	"library-app/entities"
	"net/http"
)

// SignUp godoc
// @Summary Sign up
// @Description Create new user in the system
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body entities.UserCreate true "new user info"
// @Success 200 {object} string "new user id"
// @Failure 400 {object} ResponseStruct
// @Failure 500 {object} ResponseStruct
// @Router /auth/sign-up [post]
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

// SignIn godoc
// @Summary Sign in
// @Description Auth user in the system to get token
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body entities.UserCreate true "user info"
// @Success 200 {object} string "token"
// @Failure 400 {object} ResponseStruct
// @Failure 500 {object} ResponseStruct
// @Router /auth/sign-in [post]
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
