package handler

import (
	"github.com/gin-gonic/gin"
	"library-app/entities"
	"net/http"
	"strconv"
)

// createBook godoc
// @Summary Create book
// @Description Create new book in the system
// @Security ApiKeyAuth
// @Tags Books
// @Accept json
// @Produce json
// @Param input body entities.BookCreate true "new book info"
// @Success 200 {object} string "new book id"
// @Failure 400 {object} ResponseStruct
// @Failure 500 {object} ResponseStruct
// @Router /api/books [post]
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

// getBooks godoc
// @Summary Get books
// @Description Get all the books in the system
// @Security ApiKeyAuth
// @Tags Books
// @Produce json
// @Success 200 {object} entities.Books "books data"
// @Failure 500 {object} ResponseStruct
// @Router /api/books [get]
func (h *Handler) getBooks(ctx *gin.Context) {
	books, err := h.service.Books.GetBooks()
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, books)
}

// getBookById godoc
// @Summary Get book
// @Description Get the book by id in the system
// @Security ApiKeyAuth
// @Tags Books
// @Produce json
// @Success 200 {object} entities.BookGet "book data"
// @Failure 500 {object} ResponseStruct
// @Router /api/books/:book [get]
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

// updateBookById godoc
// @Summary Update book
// @Description Update book data by id in the system
// @Security ApiKeyAuth
// @Tags Books
// @Accept json
// @Param input body entities.BookUpdate true "new genre data"
// @Success 200
// @Failure 400 {object} ResponseStruct
// @Failure 500 {object} ResponseStruct
// @Router /api/books/:book [put]
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

// deleteBookById godoc
// @Summary Delete book
// @Description Delete book by id in the system
// @Security ApiKeyAuth
// @Tags Books
// @Success 200
// @Failure 400 {object} ResponseStruct
// @Failure 500 {object} ResponseStruct
// @Router /api/books/:book [delete]
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
