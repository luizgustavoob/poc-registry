package application

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func NewServer(r http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:         ":8080",
			Handler:      r,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}

func (s *Server) ListenAndServe() {
	go func() {
		log.Printf("registry service running on %s", s.server.Addr)
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("error starting registry service: %q", err)
		}
	}()
}

func (s *Server) Shutdown() {
	log.Print("shutting down registry service...")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	err := s.server.Shutdown(ctx)
	if err != nil && err != http.ErrServerClosed {
		log.Printf("unable to shutdown the registry service in 60s: %q", err)
		return
	}

	log.Print("registry service gracefully stopped")
}
