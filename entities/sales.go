package entities

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

type SalesPayload struct {
	ID         uuid.UUID  `json:"id,omitempty"`
	ProductsID uuid.UUID  `json:"product_id"`
	CustomerID uuid.UUID  `json:"customer_id"`
	Quantity   int        `json:"quantity"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}

func (l SalesPayload) Validate() error {
	return validation.ValidateStruct(
		&l,
		validation.Field(&l.ProductsID, validation.Required),
		validation.Field(&l.CustomerID, validation.Required),
		validation.Field(&l.Quantity, validation.Required),
	)
}

type SalesResponse struct {
	ID           uuid.UUID  `json:"id"`
	ProductName  string     `json:"product_name"`
	CustomerName string     `json:"customer_name"`
	Quantity     int        `json:"quantity"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}
