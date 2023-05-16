package services

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/luizgustavoob/registry/internal/entities"
)

type (
	client interface {
		Do(req *http.Request) (*http.Response, error)
	}

	registry interface {
		ListServices() []entities.RemoteService
		DeleteService(serviceName string)
	}

	hearbeatservice struct {
		registry registry
		client   client
	}
)

func Heartbeat(registry registry, client client) *hearbeatservice {
	return &hearbeatservice{
		registry: registry,
		client:   client,
	}
}

func (h *hearbeatservice) Run() {
	go func() {
		for {
			srvs := h.registry.ListServices()

			for _, srv := range srvs {
				req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/ping", srv.Address), nil)

				resp, err := h.client.Do(req)
				if err != nil {
					log.Printf("error requesting the service %s. unregister now.", srv.ProcessName)
					h.registry.DeleteService(srv.ProcessName)
					continue
				}

				defer resp.Body.Close()

				if resp.StatusCode != http.StatusOK {
					log.Printf("service %s not ok. unregister now.", srv.ProcessName)
					h.registry.DeleteService(srv.ProcessName)
					continue
				}

				log.Printf("great! service %s ok.", srv.ProcessName)
				time.Sleep(1 * time.Second)
			}

			time.Sleep(60 * time.Second)
		}
	}()
}
