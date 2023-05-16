package database

import (
	"fmt"
	"log"

	"github.com/luizgustavoob/registry/internal/entities"
)

type InMemoryDatabase struct {
	services map[string]entities.RemoteService
}

func NewInMemoryDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{
		services: make(map[string]entities.RemoteService),
	}
}

func (db *InMemoryDatabase) AddService(srv entities.RemoteService) error {
	_, ok := db.services[srv.ProcessName]
	if ok {
		return fmt.Errorf("service %s already registered", srv.ProcessName)
	}

	db.services[srv.ProcessName] = srv
	log.Printf("ok! service %s registered", srv.ProcessName)

	return nil
}

func (db *InMemoryDatabase) DeleteService(serviceName string) {
	_, ok := db.services[serviceName]
	if !ok {
		return
	}

	delete(db.services, serviceName)
	log.Printf("service %s unregistered", serviceName)
}

func (db *InMemoryDatabase) FindService(serviceName string) (entities.RemoteService, error) {
	service, ok := db.services[serviceName]
	if !ok {
		return entities.RemoteService{}, fmt.Errorf("service %s not registered", serviceName)
	}

	return service, nil
}

func (db *InMemoryDatabase) ListServices() []entities.RemoteService {
	srvs := []entities.RemoteService{}

	for _, srv := range db.services {
		srvs = append(srvs, srv)
	}

	return srvs
}
