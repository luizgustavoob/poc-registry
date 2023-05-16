package services

import (
	"log"
)

type WorkOrder struct{}

func NewWorkOrderService() *WorkOrder {
	return &WorkOrder{}
}

func (w *WorkOrder) Create() {
	log.Print("Faz de conta que estamos criando uma work-order aqui em shipment-injection garotinho...")
}
