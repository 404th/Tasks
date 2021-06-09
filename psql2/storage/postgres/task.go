package postgres

import (
	"fmt"

	"github.com/404th/psql2/storage/models"
	"github.com/404th/psql2/storage/repo"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type taskRepo struct {
	db *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) repo.TaskStorageI {
	return &taskRepo{db: db}
}

func (tk *taskRepo) Create(m models.Task) (string, error) {
	id := uuid.New().String()

	query := `
		INSERT INTO tasks (
			id,
			title,
			body,
			done
		) VALUES (
			$1, $2, $3, $4
		)
	`
	_, err := tk.db.Exec(query,
		id,
		m.Title,
		m.Body,
		m.Done,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (tk *taskRepo) Get(id string) (models.Task, error) {
	var getTask models.Task

	row := tk.db.QueryRow(`
		SELECT 
			id,
			title,
			body,
			done FROM tasks WHERE id=$1
	`, id)
	err := row.Scan(&getTask.ID, &getTask.Title, &getTask.Body, &getTask.Done)
	if err != nil {
		return models.Task{}, err
	}

	return getTask, err
}

func (tk *taskRepo) GetAll() ([]models.Task, error) {
	var tasks []models.Task

	row := tk.db.QueryRow(`SELECT * FROM tasks`)
	err := row.Scan(&tasks)
	if err != nil {
		return []models.Task{}, err
	}

	return tasks, nil
}

func (tk *taskRepo) Update(m models.Task) (string, error) {

	query := fmt.Sprintf(`
		UPDATE tasks 
		SET 
			title=$2,
			body=$3,
			done=$4
		WHERE 
			id=$5
	`)
	_, err := tk.db.Exec(query,
		m.Title,
		m.Body,
		m.Done,
		m.ID,
	)

	if err != nil {
		return "", err
	}

	return m.ID, nil
}

func (tk *taskRepo) Delete(id string) (string, error) {
	var deletedId string

	query := `DELETE FROM tasks WHERE id = $1 RETURNING id`
	row := tk.db.QueryRow(query, id)
	err := row.Scan(&deletedId)
	if err != nil {
		return "", err
	}

	return deletedId, nil
}
