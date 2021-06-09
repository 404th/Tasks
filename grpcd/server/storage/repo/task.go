package repo

import (
	"github.com/404th/grpcd/protos"
)

type TaskStorageI interface {
	Create(task *protos.Task) (*protos.WithID, error)
	Get(task *protos.WithID) (*protos.Task, error)
	GetAll() (*protos.Tasks, error)
	Update(task *protos.Task) (*protos.WithID, error)
	Delete(withId *protos.WithID) (*protos.WithID, error)
}
