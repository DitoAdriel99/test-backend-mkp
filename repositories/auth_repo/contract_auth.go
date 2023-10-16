package auth_repo

import (
	"database/sql"
	"go-learn/config"
	"go-learn/entities"
)

type _AuthRepoImp struct {
	conn *sql.DB
}

type AuthContract interface {
	Checklogin(auth *entities.Login) (*entities.User, error)
	ValidateUser(email string) (*entities.User, error)
	CheckName(name string) error
	Register(rq *entities.Register) error
}

func NewAuthRepositories() AuthContract {
	conn, err := config.DBConn()
	if err != nil {
		panic(err)
	}

	return &_AuthRepoImp{
		conn: conn,
	}
}
