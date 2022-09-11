package handler

import (
	"github.com/gin-gonic/gin"
	"library-app/entities"
	"net/http"
	"strconv"
)

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
