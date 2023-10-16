package product_service

import (
	"fmt"
	"go-learn/entities"
	"time"

	"github.com/google/uuid"
)

func (s *_Service) CreateSales(payload entities.SalesPayload) error {
	var (
		time   = time.Now().Local()
		id, _  = uuid.NewUUID()
		reqQty int
	)
	payload.ID = id
	payload.CreatedAt = time
	payload.UpdatedAt = &time

	_, err := s.repo.CustomerRepo.Detail(payload.CustomerID)
	if err != nil {
		return err
	}

	prod, err := s.repo.ProductRepo.Detail(payload.ProductsID)
	if err != nil {
		return err
	}

	if prod.Stock < payload.Quantity {
		return fmt.Errorf("Out Of Stock")
	}

	reqQty = prod.Stock - int(payload.Quantity)

	if err := s.repo.ProductRepo.CreateSales(payload, reqQty); err != nil {
		return err
	}
	return nil
}
