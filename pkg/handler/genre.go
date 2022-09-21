package handler

import (
	"github.com/gin-gonic/gin"
	"library-app/entities"
	"net/http"
	"strconv"
)

// createGenre godoc
// @Summary Create genre
// @Description Create new genre in the system
// @Security ApiKeyAuth
// @Tags Genres
// @Accept json
// @Produce json
// @Param input body entities.GenreCreateAndGet true "new genre info"
// @Success 200 {object} string "new genre id"
// @Failure 400 {object} ResponseStruct
// @Failure 500 {object} ResponseStruct
// @Router /api/genres [post]
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

// getGenres godoc
// @Summary Get genres
// @Description Get all the genres in the system
// @Security ApiKeyAuth
// @Tags Genres
// @Produce json
// @Success 200 {object} entities.Genres "genres data"
// @Failure 500 {object} ResponseStruct
// @Router /api/genres [get]
func (h *Handler) getGenres(ctx *gin.Context) {
	genres, err := h.service.Genres.GetGenres()
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, genres)
}

// getGenreById godoc
// @Summary Get genre
// @Description Get the genre by id in the system
// @Security ApiKeyAuth
// @Tags Genres
// @Produce json
// @Success 200 {object} entities.GenreCreateAndGet "genre data"
// @Failure 500 {object} ResponseStruct
// @Router /api/genres/:genre [get]
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

// @Summary Update genre
// @Description Update genre data by id in the system
// @Security ApiKeyAuth
// @Tags Genres
// @Accept json
// @Param input body entities.GenreUpdate true "new genre data"
// @Success 200
// @Failure 400 {object} ResponseStruct
// @Failure 500 {object} ResponseStruct
// @Router /api/genres/:genre [put]

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

// @Summary Delete genre
// @Description Delete genre by id in the system
// @Security ApiKeyAuth
// @Tags Genres
// @Success 200
// @Failure 400 {object} ResponseStruct
// @Failure 500 {object} ResponseStruct
// @Router /api/genres/:genre [delete]

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
