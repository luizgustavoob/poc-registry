package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func BuildRoutes(app *App) http.Handler {
	r := chi.NewRouter()

	//middlewares
	r.Use(remoteServiceMiddleware(app.registry))

	//registry
	r.Get("/services", app.HandleListServices)
	r.Post("/services", app.HandleAddService)
	r.Delete("/services/{service_name}", app.HandleDeleteService)

	//proxy
	r.HandleFunc("/commands", proxy)

	return r
}
