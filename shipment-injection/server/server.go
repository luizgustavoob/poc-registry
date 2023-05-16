package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/luizgustavoob/shipment-injection/internal/core/consts"
	"github.com/luizgustavoob/shipment-injection/internal/core/ports"
)

type Handler interface {
	Method() string
	Pattern() string
	http.Handler
}

type Server struct {
	registry ports.Registry
	server   *http.Server
}

func NewServer(registry ports.Registry, handlers ...Handler) *Server {
	r := chi.NewRouter()
	for _, h := range handlers {
		r.Method(h.Method(), h.Pattern(), h)
	}

	return &Server{
		registry: registry,
		server: &http.Server{
			Addr:         ":" + consts.ServerPort,
			Handler:      r,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}

func (s *Server) ListenAndServe() {
	go func() {
		err := s.registry.Register()
		if err != nil {
			log.Printf("error registering service: %s", err.Error())
			panic(err)
		}

		log.Printf("%s running on %s", consts.ProcessName, s.server.Addr)

		err = s.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Printf("error starting %s: %q", consts.ProcessName, err)
			panic(err)
		}
	}()
}

func (s *Server) Shutdown() {
	log.Printf("shutting down %s", consts.ProcessName)

	s.registry.Unregister()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		log.Printf("unable to shutdown the %s in 60s: %q", consts.ProcessName, err)
		return
	}

	log.Printf("%s gracefully stopped", consts.ProcessName)
}
