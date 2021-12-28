package service

import (
	"todolist/model"
	"todolist/repository"
)

type ToDoList interface {
	Create(model model.TodoListInput) (model.ToDoList, error)
}

type service struct {
	repo repository.ToDoList
}

func NewService(repo repository.ToDoList) *service {
	return &service{repo}
}

func (s *service) Create(input model.TodoListInput) (model.ToDoList, error) {

	var newModel = model.ToDoList{
		Task:   input.Task,
		Status: false,
	}
	dt, err := s.repo.Create(newModel)

	if err != nil {
		return dt, nil
	}

	return dt, nil
}
