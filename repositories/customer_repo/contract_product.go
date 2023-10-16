package customer_repo

import (
	"database/sql"
	"go-learn/config"
	"go-learn/entities"

	"github.com/google/uuid"
)

type _CustomerRepoImp struct {
	conn *sql.DB
}

type CustomerContract interface {
	Detail(id uuid.UUID) (*entities.Customer, error)
}

func NewCustomerRepositories() CustomerContract {
	conn, err := config.DBConn()
	if err != nil {
		panic(err)
	}

	return &_CustomerRepoImp{
		conn: conn,
	}
}
