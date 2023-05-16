package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/luizgustavoob/hu-assembly/internal/adapters/handlers"
	"github.com/luizgustavoob/hu-assembly/internal/services"
	"github.com/luizgustavoob/hu-assembly/server"
)

func main() {
	registry := services.NewRegistry(&http.Client{Timeout: 500 * time.Millisecond})
	workOrderService := services.NewWorkOrderService()

	// handlers
	pingHandler := handlers.NewPingHandler()
	createWorkOrderHandler := handlers.NewCreateWorkOrderHandler(workOrderService)

	// server
	srv := server.NewServer(registry, pingHandler, createWorkOrderHandler)
	srv.ListenAndServe()

	// shutdown
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan
	srv.Shutdown()
}
