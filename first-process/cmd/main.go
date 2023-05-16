package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/luizgustavoob/first-process/internal/adapters/handlers"
	"github.com/luizgustavoob/first-process/internal/services"
	"github.com/luizgustavoob/first-process/server"
)

func main() {
	registry := services.NewRegistry(&http.Client{Timeout: 500 * time.Millisecond})
	orderService := services.NewOrderService()

	// handlers
	pingHandler := handlers.NewPingHandler()
	createOrderHandler := handlers.NewCreateOrderHandler(orderService)

	// server
	srv := server.NewServer(registry, pingHandler, createOrderHandler)
	srv.ListenAndServe()

	// shutdown
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan
	srv.Shutdown()
}
