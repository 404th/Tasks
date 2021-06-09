package repo

import "github.com/404th/Tasks/psql/storage/models"

type ContactStorageI interface {
	Create(m models.Contact) (string, error)
	Get(id string) (models.Contact, error)
	GetAll() ([]models.Contact, error)
	Update(m models.Contact) error
	Delete(id string) (string, error)
}
