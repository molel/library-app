package handler

import (
	"github.com/gin-gonic/gin"
	"library-app/entities"
	"net/http"
	"strconv"
)

// createListItem godoc
// @Summary Create list item
// @Description Create new list item in the system
// @Security ApiKeyAuth
// @Tags Items
// @Accept json
// @Param input body entities.ListItemCreate true "new list item info"
// @Success 200
// @Failure 400 {object} ResponseStruct
// @Failure 500 {object} ResponseStruct
// @Router /api/lists/:list/items [post]
func (h *Handler) createListItem(ctx *gin.Context) {
	var inputJSON entities.ListItemCreate
	if err := ctx.BindJSON(&inputJSON); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	userId, err := GetUserId(ctx)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	listId, err := strconv.Atoi(ctx.Param("list"))
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	err = h.service.ListItems.CreateListItem(userId, listId, inputJSON)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusOK)
}

// updateListItemById godoc
// @Summary Update list item
// @Description Update list item data by id in the system
// @Security ApiKeyAuth
// @Tags Items
// @Accept json
// @Param input body entities.ListItemUpdate true "new genre data"
// @Success 200
// @Failure 400 {object} ResponseStruct
// @Failure 500 {object} ResponseStruct
// @Router /api/lists/:list/items/:book [put]
func (h *Handler) updateListItemById(ctx *gin.Context) {
	var inputJSON entities.ListItemUpdate
	if err := ctx.BindJSON(&inputJSON); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	userId, err := GetUserId(ctx)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	listId, err := strconv.Atoi(ctx.Param("list"))
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	bookId, err := strconv.Atoi(ctx.Param("book"))
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	err = h.service.ListItems.UpdateListItemById(userId, listId, bookId, inputJSON)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusOK)
}

// deleteListItemById godoc
// @Summary Delete list item
// @Description Delete list item by id in the system
// @Security ApiKeyAuth
// @Tags Items
// @Success 200
// @Failure 400 {object} ResponseStruct
// @Failure 500 {object} ResponseStruct
// @Router /api/lists/:list/items/:book [delete]
func (h *Handler) deleteListItemById(ctx *gin.Context) {
	userId, err := GetUserId(ctx)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	listId, err := strconv.Atoi(ctx.Param("list"))
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	bookId, err := strconv.Atoi(ctx.Param("book"))
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	err = h.service.ListItems.DeleteListItemById(userId, listId, bookId)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusOK)
}
