package models

type User struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Age int32 `json:"age"`
	IsAdult bool `json:"is_adult"`
}