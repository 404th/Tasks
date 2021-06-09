package service

import "github.com/404th/contactlist/pkg/repo"

type Authorization interface{}

type TodoLists interface{}

type TodoItem interface{}

type Service struct {
	Authorization
	TodoLists
	TodoItem
}

func NewService(repos *repo.Repo) *Service {
	return &Service{}
}
