package product_service

import (
	"go-learn/entities"
	"go-learn/library/meta"
	"log"
)

func (s *_Service) GetSales(m *meta.Metadata) ([]entities.SalesResponse, error) {
	resp, err := s.repo.ProductRepo.GetSales(m)
	if err != nil {
		log.Println("Get All is error : ", err)
		return nil, err
	}

	return resp, nil
}
