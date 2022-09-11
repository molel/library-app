package service

import (
	"library-app/entities"
	"library-app/pkg/repository"
)

type ListItemService struct {
	repository *repository.Repository
}

func NewListItemService(repository *repository.Repository) *ListItemService {
	return &ListItemService{repository: repository}
}

func (lis *ListItemService) CreateListItem(userId, listId int, listItem entities.ListItemCreate) error {
	return lis.repository.CreateListItem(userId, listId, listItem)
}

func (lis *ListItemService) UpdateListItemById(userId, listId, bookId int, listItem entities.ListItemUpdate) error {
	return lis.repository.UpdateListItemById(userId, listId, bookId, listItem)
}

func (lis *ListItemService) DeleteListItemById(userId, listId, bookId int) error {
	return lis.repository.ListItems.DeleteListItemById(userId, listId, bookId)
}
