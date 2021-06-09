package storage

import (
	"github.com/404th/psql2/storage/postgres"
	"github.com/404th/psql2/storage/repo"
	"github.com/jmoiron/sqlx"
)

type storagePg struct {
	db       *sqlx.DB
	taskRepo repo.TaskStorageI
}

func (s storagePg) Task() repo.TaskStorageI {
	return s.taskRepo
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		db:       db,
		taskRepo: postgres.NewTaskRepo(db),
	}
}

type StorageI interface {
	Task() repo.TaskStorageI
}
