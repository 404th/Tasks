package postgres

import (
	"log"

	"github.com/404th/grpce/models"
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sqlx.DB
}

func NewRepo (db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Create(id, username string, age int32, is_adult bool) (string, bool, error) {

	query := `INSERT INTO users ( id, username, age, is_adult ) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, id, username, age, is_adult)
	if err != nil {
		log.Printf("error while creating user: %v", err)
		return "", false, err
	}
	return id, true, nil
}

func (r *Repo) Get(id string) ( *models.User, error) {
	var usr *models.User

	query := `SELECT id, username, age, is_adult FROM users WHERE id=$1`
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&usr.ID, &usr.Username, &usr.Age, &usr.IsAdult); err != nil {
		log.Printf("error while creating user: %v", err)
		return nil, err
	}

	return usr, nil
}

func (r *Repo) Delete(id string) (string, bool, error) {

	query := `DELETE FROM users WHERE id=$1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("error while deleting user: %v", err)
		return "", false, err
	}
	return id, true, nil
}

