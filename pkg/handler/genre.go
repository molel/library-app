package handler

import (
	"github.com/gin-gonic/gin"
	"library-app/entities"
	"net/http"
	"strconv"
)

func (h *Handler) createGenre(ctx *gin.Context) {
	var inputJSON entities.GenreCreateAndGet
	if err := ctx.BindJSON(&inputJSON); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	genreId, err := h.service.Genres.CreateGenre(inputJSON)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"id": genreId})
}

func (h *Handler) getGenres(ctx *gin.Context) {
	genres, err := h.service.Genres.GetGenres()
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, genres)
}

func (h *Handler) getGenreById(ctx *gin.Context) {
	intId, err := strconv.Atoi(ctx.Param("genre"))
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	genre, err := h.service.Genres.GetGenreById(intId)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, genre)
}

func (h *Handler) updateGenreById(ctx *gin.Context) {
	var inputJSON entities.GenreUpdate
	if err := ctx.BindJSON(&inputJSON); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	intId, err := strconv.Atoi(ctx.Param("genre"))
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	err = h.service.Genres.UpdateGenreById(intId, inputJSON)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusOK)
}

func (h *Handler) deleteGenreById(ctx *gin.Context) {
	intId, err := strconv.Atoi(ctx.Param("genre"))
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	err = h.service.Genres.DeleteGenreById(intId)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusOK)
}
