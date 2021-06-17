package repo

import "github.com/404th/grpce/models"
type UserStorageI interface {
	Create(id, username string, age int32, is_adult bool) (string, bool, error)
	Get(id string) ( *models.User, error)
	Delete(id string) (string, bool, error)
}
