package entities

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID           uuid.UUID  `json:"id"`
	CustomerName string     `json:"customer_name"`
	ContactInfo  string     `json:"contact_info"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}
