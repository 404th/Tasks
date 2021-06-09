package models

type Contact struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Username    string `db:"username"`
	PhoneNumber string `db:"phone_number"`
	Profession  string `db:"profession"`
	Age         int    `db:"age"`
	IsAdmin     bool   `db:"is_admin"`
}
