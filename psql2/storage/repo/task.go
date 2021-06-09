package repo

import "github.com/404th/psql2/storage/models"

type TaskStorageI interface {
	Create(m models.Task) (string, error)
	Get(id string) (models.Task, error)
	GetAll() ([]models.Task, error)
	Update(m models.Task) (string, error)
	Delete(id string) (string, error)
}
