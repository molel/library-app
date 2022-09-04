package handler

import (
	"github.com/gin-gonic/gin"
	"library-app/entities"
	"net/http"
)

func (h *Handler) createAuthor(ctx *gin.Context) {
	var inputJSON entities.Author
	if err := ctx.BindJSON(&inputJSON); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	authorId, err := h.service.Authors.CreateAuthor(inputJSON)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"authorId": authorId})
}

func (h *Handler) getAuthors(ctx *gin.Context) {
	authors, err := h.service.Authors.GetAuthors()
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, authors)
}
