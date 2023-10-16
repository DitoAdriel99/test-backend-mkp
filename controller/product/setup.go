package product

import "go-learn/service"

type _ControllerProduct struct {
	service service.Service
}

func NewControllerProductCreate(service service.Service) *_ControllerProduct {
	return &_ControllerProduct{service: service}
}
