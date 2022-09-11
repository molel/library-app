package service

import (
	"library-app/entities"
	"library-app/pkg/repository"
)

type ListService struct {
	repository *repository.Repository
}

func NewListService(repository *repository.Repository) *ListService {
	return &ListService{repository: repository}
}

func (ls *ListService) CreateList(userId int, list entities.ListCreate) (int, error) {
	return ls.repository.Lists.CreateList(userId, list)
}

func (ls *ListService) GetLists(userId int) ([]entities.ListGet, error) {
	return ls.repository.Lists.GetLists(userId)
}

func (ls *ListService) GetListById(userId, id int) (entities.ListGet, error) {
	return ls.repository.Lists.GetListById(userId, id)
}

func (ls *ListService) UpdateListById(userId, id int, list entities.ListUpdate) error {
	return ls.repository.Lists.UpdateListById(userId, id, list)
}

func (ls *ListService) DeleteListById(userId, id int) error {
	return ls.repository.Lists.DeleteListById(userId, id)
}
