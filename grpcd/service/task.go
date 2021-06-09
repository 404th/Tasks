package service

import (
	"context"
	"log"

	"github.com/404th/grpcd/protos"
	"github.com/404th/grpcd/server/storage"
	"github.com/google/uuid"
)

type TaskService struct {
	storage storage.StorageI
}

func NewTaskService(strg storage.StorageI) *TaskService {
	return &TaskService{
		storage: strg,
	}
}

func (s *TaskService) Create(ctx context.Context, req *protos.Task) (*protos.WithID, error) {
	req.Id = uuid.New().String()
	withId, err := s.storage.Task().Create(req)
	if err != nil {
		log.Printf("error while creating: %v", err)
		return nil, err
	}
	return &protos.WithID{
		Id: withId.GetId(),
	}, nil
}

func (s *TaskService) Get(ctx context.Context, req *protos.WithID) (*protos.Task, error) {
	res, err := s.storage.Task().Get(req)
	if err != nil {
		log.Printf("error while getting: %v", err)
		return nil, err
	}
	return res, nil
}

func (s *TaskService) GetAll(ctx context.Context, req *protos.Empty) (*protos.Tasks, error) {
	res, err := s.storage.Task().GetAll()
	if err != nil {
		log.Printf("error while getting all: %v", err)
		return nil, err
	}
	return res, nil
}

func (s *TaskService) Update(ctx context.Context, req *protos.Task) (*protos.WithID, error) {
	res, err := s.storage.Task().Update(req)
	if err != nil {
		log.Printf("error while updating: %v", err)
		return nil, err
	}
	return res, nil
}

func (s *TaskService) Delete(ctx context.Context, req *protos.WithID) (*protos.WithID, error) {
	res, err := s.storage.Task().Delete(req)
	if err != nil {
		log.Printf("error while updating: %v", err)
		return nil, err
	}
	return res, nil
}
