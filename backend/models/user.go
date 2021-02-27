package models

type User struct {
	ID       int    `json:"id"`
	UserID   string `json:"userId"`
	Password string `json:"password"`
}
