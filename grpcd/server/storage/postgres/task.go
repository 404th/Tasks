package postgres

import (
	"fmt"

	"github.com/404th/grpcd/protos"
	"github.com/404th/grpcd/server/storage/repo"
	"github.com/jmoiron/sqlx"
)

type taskRepo struct {
	db *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) repo.TaskStorageI {
	return &taskRepo{
		db: db,
	}
}

func (s *taskRepo) Create(task *protos.Task) (*protos.WithID, error) {
	query := `
		INSERT INTO tasks 
		(
			id, 
			title, 
			body
		) 
			VALUES 
		(
			$1, 
			$2, 
			$3
		)
	`

	_, err := s.db.Exec(query,
		task.GetId(),
		task.GetTitle(),
		task.GetBody(),
	)

	if err != nil {
		fmt.Printf("error while creating: %v", err)
		return nil, err
	}

	return &protos.WithID{
		Id: task.GetId(),
	}, nil
}

func (s *taskRepo) Get(task *protos.WithID) (*protos.Task, error) {
	var tsk protos.Task
	query := `
		SELECT 
			title, 
			body 
		FROM 
			tasks 
		WHERE 
			id=$1
		`

	row := s.db.QueryRow(query,
		task.Id,
	)

	if err := row.Scan(&tsk.Title, &tsk.Body); err != nil {
		fmt.Printf("error while getting: %v", err)
		return nil, err
	}
	tsk.Id = task.GetId()

	return &tsk, nil
}

func (s *taskRepo) GetAll() (*protos.Tasks, error) {
	var tsks protos.Tasks
	query := `SELECT * FROM tasks`

	rows, err := s.db.Query(query)
	if err != nil {
		fmt.Printf("error while getting all: %v", err)
		return nil, err
	}

	for rows.Next() {
		var tsk protos.Task

		if err := rows.Scan(
			&tsk.Id,
			&tsk.Title,
			&tsk.Body,
		); err != nil {
			fmt.Printf("error while getting all: %v", err)
			return nil, err
		}

		tsks.Tasks = append(tsks.Tasks, &tsk)
	}

	return &tsks, nil
}

func (s *taskRepo) Update(task *protos.Task) (*protos.WithID, error) {
	query :=
		`UPDATE 
		tasks 
	SET 
		title=$1, 
		body=$2 
	WHERE 
		id=$3
	`

	_, err := s.db.Exec(
		query,
		task.GetTitle(),
		task.GetBody(),
		task.GetId(),
	)

	if err != nil {
		fmt.Printf("error while updating: %v", err)
		return nil, err
	}

	return &protos.WithID{
		Id: task.GetId(),
	}, nil
}

func (s *taskRepo) Delete(withId *protos.WithID) (*protos.WithID, error) {
	query := `DELETE FROM tasks WHERE id=$1`

	_, err := s.db.Exec(
		query,
		withId.GetId(),
	)

	if err != nil {
		fmt.Printf("error while deleting: %v", err)
		return nil, err
	}

	return &protos.WithID{
		Id: withId.GetId(),
	}, nil
}
