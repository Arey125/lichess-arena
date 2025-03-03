package users

import (
	"database/sql"
)

type User struct {
}

type UserModel struct {
	db *sql.DB
}

func NewModel(db *sql.DB) UserModel {
	return UserModel{db}
}
