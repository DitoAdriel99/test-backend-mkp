package product_service

import (
	"go-learn/entities"
	"go-learn/library/meta"
	"go-learn/repositories"

	"github.com/google/uuid"
)

type Contract interface {
	Create(payload *entities.Product) error
	Detail(id uuid.UUID) (*entities.Product, error)
	GetAll(m *meta.Metadata) ([]entities.Product, error)
	CreateSales(payload entities.SalesPayload) error
	GetSales(m *meta.Metadata) ([]entities.SalesResponse, error)
}

type _Service struct {
	repo *repositories.Repo
}

func NewProductService(repo *repositories.Repo) Contract {
	return &_Service{repo}
}
