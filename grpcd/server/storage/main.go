package storage

import (
	"github.com/404th/grpcd/server/storage/postgres"
	"github.com/404th/grpcd/server/storage/repo"
	"github.com/jmoiron/sqlx"
)

type storagePg struct {
	taskRepo repo.TaskStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		taskRepo: postgres.NewTaskRepo(db),
	}
}

func (s storagePg) Task() repo.TaskStorageI {
	return s.taskRepo
}

type StorageI interface {
	Task() repo.TaskStorageI
}
