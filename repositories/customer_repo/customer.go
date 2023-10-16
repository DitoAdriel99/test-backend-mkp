package customer_repo

import (
	"database/sql"
	"fmt"
	"go-learn/entities"

	"github.com/google/uuid"
)

func (c *_CustomerRepoImp) Detail(id uuid.UUID) (*entities.Customer, error) {
	query := `
		SELECT 
			* 
		FROM 
			customers 
		WHERE 
			id = $1`

	var object entities.Customer

	err := c.conn.QueryRow(query, id).Scan(
		&object.ID,
		&object.CustomerName,
		&object.ContactInfo,
		&object.CreatedAt,
		&object.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("customer not found")
		}
		err = fmt.Errorf("scanning activity objects: %w", err)
		return nil, err
	}

	return &object, nil
}
