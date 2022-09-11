package handler

import (
	"github.com/gin-gonic/gin"
	"library-app/entities"
	"net/http"
	"strconv"
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
	ctx.JSON(http.StatusOK, map[string]interface{}{"id": bookId})
}

func (h *Handler) getBooks(ctx *gin.Context) {
	books, err := h.service.Books.GetBooks()
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, books)
}

func (h *Handler) getBookById(ctx *gin.Context) {
	intId, err := strconv.Atoi(ctx.Param("book"))
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	book, err := h.service.Books.GetBookById(intId)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func (h *Handler) updateBookById(ctx *gin.Context) {
	var inputJSON entities.BookUpdate
	if err := ctx.BindJSON(&inputJSON); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	intId, err := strconv.Atoi(ctx.Param("book"))
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	err = h.service.UpdateBookById(intId, inputJSON)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusOK)
}

func (h *Handler) deleteBookById(ctx *gin.Context) {
	intId, err := strconv.Atoi(ctx.Param("book"))
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	err = h.service.Books.DeleteBookById(intId)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusOK)
}
