package handler

import (
	"github.com/gin-gonic/gin"
	"library-app/entities"
	"net/http"
)

func (h *Handler) createBook(ctx *gin.Context) {
	var inputJSON entities.BookCreate
	if err := ctx.BindJSON(&inputJSON); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	bookId, err := h.service.Books.CreateBook(inputJSON)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"bookId": bookId})
}
