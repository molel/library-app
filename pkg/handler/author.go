package handler

import (
	"github.com/gin-gonic/gin"
	"library-app/entities"
	"net/http"
	"strconv"
)

// @Summary Create author
// @Description Create new author in the system
// @Security ApiKeyAuth
// @Tags Authors
// @Accept json
// @Produce json
// @Param input body entities.AuthorCreate true "new author info"
// @Success 200 {object} string "new author id"
// @Failure 400 {object} ResponseStruct
// @Failure 500 {object} ResponseStruct
// @Router /api/authors [post]
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

// @Summary Get authors
// @Description Get all the authors in the system
// @Security ApiKeyAuth
// @Tags Authors
// @Produce json
// @Success 200 {object} entities.Authors "authors data"
// @Failure 500 {object} ResponseStruct
// @Router /api/authors [get]
func (h *Handler) getAuthors(ctx *gin.Context) {
	authors, err := h.service.Authors.GetAuthors()
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, authors)
}

// @Summary Get author
// @Description Get the author by id in the system
// @Security ApiKeyAuth
// @Tags Authors
// @Produce json
// @Success 200 {object} entities.AuthorGet "author data"
// @Failure 500 {object} ResponseStruct
// @Router /api/authors/:author [get]
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

// @Summary Update author
// @Description Update author data by id in the system
// @Security ApiKeyAuth
// @Tags Authors
// @Accept json
// @Param input body entities.AuthorUpdate true "new author data"
// @Success 200
// @Failure 400 {object} ResponseStruct
// @Failure 500 {object} ResponseStruct
// @Router /api/authors/:author [put]
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

// @Summary Delete author
// @Description Delete author by id in the system
// @Security ApiKeyAuth
// @Tags Authors
// @Success 200
// @Failure 400 {object} ResponseStruct
// @Failure 500 {object} ResponseStruct
// @Router /api/authors/:author [delete]
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
