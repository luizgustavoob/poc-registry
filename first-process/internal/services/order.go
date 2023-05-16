package services

import (
	"log"
)

type orderService struct{}

func NewOrderService() *orderService {
	return &orderService{}
}

func (*orderService) Create() {
	log.Print("first-process order...")
}
