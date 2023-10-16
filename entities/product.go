package entities

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

type Product struct {
	ID           uuid.UUID  `json:"id"`
	ProductName  string     `json:"product_name"`
	ProductPrice int        `json:"product_price"`
	Stock        int        `json:"stock"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

func (l Product) Validate() error {
	return validation.ValidateStruct(
		&l,
		validation.Field(&l.ProductName, validation.Required),
		validation.Field(&l.ProductPrice, validation.Required),
		validation.Field(&l.Stock, validation.Required),
	)
}
