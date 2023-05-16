package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/luizgustavoob/registry/infrastructure/database"
	"github.com/luizgustavoob/registry/internal/application"
	"github.com/luizgustavoob/registry/internal/services"
)

func main() {
	db := database.InMemory()

	app := application.BuildApp(db)
	routes := application.BuildRoutes(app)
	srv := application.NewServer(routes)

	//run
	srv.ListenAndServe()

	//hearbeat
	services.Heartbeat(db, &http.Client{Timeout: 800 * time.Millisecond}).Run()

	//shutdown
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan
	srv.Shutdown()
}
