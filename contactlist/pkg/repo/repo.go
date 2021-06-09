package repo

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type TodoLists interface {
}

type TodoItem interface {
}

type Repo struct {
	Authorization
	TodoLists
	TodoItem
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{}
}
