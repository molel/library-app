package handler

import (
	"github.com/gin-gonic/gin"
	"library-app/entities"
	"net/http"
	"strconv"
)

// createList godoc
// @Summary Create user's book list
// @Description Create new user's book list in the system
// @Security ApiKeyAuth
// @Tags Lists
// @Accept json
// @Produce json
// @Param input body entities.ListCreate true "new user's book list info"
// @Success 200 {object} string "new user's book list id"
// @Failure 400 {object} ResponseStruct
// @Failure 500 {object} ResponseStruct
// @Router /api/lists [post]
func (h *Handler) createList(ctx *gin.Context) {
	var inputJSON entities.ListCreate
	if err := ctx.BindJSON(&inputJSON); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	userId, err := GetUserId(ctx)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	listId, err := h.service.Lists.CreateList(userId, inputJSON)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"id": listId})
}

// getLists godoc
// @Summary Get user's book lists
// @Description Get all the users' book lists in the system
// @Security ApiKeyAuth
// @Tags Lists
// @Produce json
// @Success 200 {object} entities.Lists "users' book lists data"
// @Failure 500 {object} ResponseStruct
// @Router /api/lists [get]
func (h *Handler) getLists(ctx *gin.Context) {
	userId, err := GetUserId(ctx)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	lists, err := h.service.Lists.GetLists(userId)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, lists)
}

// getListById godoc
// @Summary Get user's book list
// @Description Get the user's book list by id in the system
// @Security ApiKeyAuth
// @Tags Lists
// @Produce json
// @Success 200 {object} entities.ListGetWithItems "user's book list data"
// @Failure 500 {object} ResponseStruct
// @Router /api/lists/:list [get]
func (h *Handler) getListById(ctx *gin.Context) {
	userId, err := GetUserId(ctx)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	intId, err := strconv.Atoi(ctx.Param("list"))
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	list, err := h.service.Lists.GetListById(userId, intId)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, list)
}

// updateListById godoc
// @Summary Update user's book list
// @Description Update user's book list data by id in the system
// @Security ApiKeyAuth
// @Tags Lists
// @Accept json
// @Param input body entities.ListUpdate true "new user's book list data"
// @Success 200
// @Failure 400 {object} ResponseStruct
// @Failure 500 {object} ResponseStruct
// @Router /api/lists/:list [put]
func (h *Handler) updateListById(ctx *gin.Context) {
	var inputJSON entities.ListUpdate
	if err := ctx.BindJSON(&inputJSON); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	userId, err := GetUserId(ctx)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	intId, err := strconv.Atoi(ctx.Param("list"))
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	err = h.service.Lists.UpdateListById(userId, intId, inputJSON)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusOK)
}

// deleteListById godoc
// @Summary Delete user's book list
// @Description Delete user's book list by id in the system
// @Security ApiKeyAuth
// @Tags Lists
// @Success 200
// @Failure 400 {object} ResponseStruct
// @Failure 500 {object} ResponseStruct
// @Router /api/lists/:list [delete]
func (h *Handler) deleteListById(ctx *gin.Context) {
	userId, err := GetUserId(ctx)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	intId, err := strconv.Atoi(ctx.Param("list"))
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	err = h.service.Lists.DeleteListById(userId, intId)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusOK)
}
