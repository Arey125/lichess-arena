package users

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
    Id int
    Name string
    passHash string
}

type UserModel struct {
	db *sql.DB
}

func NewModel(db *sql.DB) UserModel {
	return UserModel{db}
}

func (m UserModel) Insert(name string, password string) error {
    passHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
    if err != nil {
        return err
    }

    q := sq.Insert("users").Columns("name", "pass_hash").Values(name, string(passHash))
    fmt.Println(q.ToSql())
    return nil
}
