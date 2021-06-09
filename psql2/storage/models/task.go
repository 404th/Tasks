package models

type Task struct {
	ID    string `db:"id"`
	Title string `db:"title"`
	Body  string `db:"body"`
	Done  bool   `db:"done"`
}
