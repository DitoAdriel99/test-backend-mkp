package entities

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
type Register struct {
	ID        uuid.UUID  `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type RegisterPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l RegisterPayload) Validate() error {
	return validation.ValidateStruct(
		&l,
		validation.Field(&l.Username, validation.Required),
		validation.Field(&l.Password, validation.Required),
	)
}
