package handler

import (
	"github.com/gin-gonic/gin"
	"library-app/entities"
	"net/http"
	"strconv"
)

func (h *Handler) createAuthor(ctx *gin.Context) {
	var inputJSON entities.AuthorCreate
	if err := ctx.BindJSON(&inputJSON); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	authorId, err := h.service.Authors.CreateAuthor(inputJSON)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"id": authorId})
}

func (h *Handler) getAuthors(ctx *gin.Context) {
	authors, err := h.service.Authors.GetAuthors()
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, authors)
}

func (h *Handler) getAuthorById(ctx *gin.Context) {
	intId, err := strconv.Atoi(ctx.Param("author"))
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	author, err := h.service.Authors.GetAuthorById(intId)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, author)
}

func (h *Handler) updateAuthorById(ctx *gin.Context) {
	var inputJSON entities.AuthorUpdate
	if err := ctx.BindJSON(&inputJSON); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	intId, err := strconv.Atoi(ctx.Param("author"))
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	err = h.service.Authors.UpdateAuthorById(intId, inputJSON)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusOK)
}

func (h *Handler) deleteAuthorById(ctx *gin.Context) {
	intId, err := strconv.Atoi(ctx.Param("author"))
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	err = h.service.Authors.DeleteAuthorById(intId)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusOK)
}
