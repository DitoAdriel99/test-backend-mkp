package auth_repo

import (
	"database/sql"
	"fmt"
	"go-learn/entities"

	"github.com/google/uuid"
)

func (c *_AuthRepoImp) Checklogin(auth *entities.Login) (*entities.User, error) {
	query := `SELECT * FROM users WHERE username = $1`

	var object entities.User

	err := c.conn.QueryRow(query, auth.Username).Scan(
		&object.ID,
		&object.Username,
		&object.Password,
		&object.CreatedAt,
		&object.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		err = fmt.Errorf("scanning activity objects: %w", err)
		return nil, err
	}

	return &object, nil

}

func (c *_AuthRepoImp) ValidateUser(email string) (*entities.User, error) {
	query := `SELECT * FROM users WHERE email = $1`

	var object entities.User

	err := c.conn.QueryRow(query, email).Scan(
		&object.ID,
		&object.Username,
		&object.CreatedAt,
		&object.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		err = fmt.Errorf("scanning activity objects: %w", err)
		return nil, err
	}
	return &object, nil
}

func (c *_AuthRepoImp) CheckName(name string) error {
	query := `SELECT COUNT(*) FROM users WHERE username = $1`

	var count int
	err := c.conn.QueryRow(query, name).Scan(&count)
	if err != nil {
		err = fmt.Errorf("scanning activity objects: %w", err)
		return err
	}

	if count == 1 {
		err = fmt.Errorf("Username Already Used!")
		return err
	}

	return nil
}

func (c *_AuthRepoImp) Register(rq *entities.Register) error {
	queryInsert := `
		INSERT INTO users (id, username, password, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
	`
	if _, err := c.conn.Exec(queryInsert, rq.ID, rq.Username, rq.Password, rq.CreatedAt, rq.UpdatedAt); err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return err
	}

	return nil
}

func (c *_AuthRepoImp) UpdateStatusUser(id uuid.UUID, status bool) error {
	updateQuery := `
	UPDATE users
	SET
		is_active = $2
	WHERE
		id = $1
	`

	if _, err := c.conn.Exec(updateQuery, id, status); err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return err
	}
	return nil
}
