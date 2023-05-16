package application

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/luizgustavoob/registry/internal/entities"
)

type (
	registry interface {
		AddService(dto entities.RemoteService) error
		DeleteService(serviceName string)
		ListServices() []entities.RemoteService
		FindService(serviceName string) (entities.RemoteService, error)
	}

	App struct {
		registry
	}
)

func BuildApp(r registry) *App {
	return &App{
		registry: r,
	}
}

func internalHandleError(w http.ResponseWriter, err error, status int) {
	result := make(map[string]string)
	result["error"] = err.Error()
	js, _ := json.Marshal(result)
	w.WriteHeader(status)
	w.Write(js)
}

func (a *App) HandleAddService(w http.ResponseWriter, r *http.Request) {
	var addr entities.RemoteService

	err := json.NewDecoder(r.Body).Decode(&addr)
	if err != nil {
		internalHandleError(w, err, http.StatusBadRequest)
		return
	}

	addr.FormatFinalAddress()
	err = a.AddService(addr)
	if err != nil {
		internalHandleError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (a *App) HandleDeleteService(w http.ResponseWriter, r *http.Request) {
	serviceName := chi.URLParam(r, "service_name")
	if serviceName == "" {
		internalHandleError(w, errors.New("service name not informed"), http.StatusBadRequest)
		return
	}

	a.DeleteService(serviceName)

	w.WriteHeader(http.StatusNoContent)
}

func (a *App) HandleListServices(w http.ResponseWriter, r *http.Request) {
	srvs := a.ListServices()
	js, _ := json.Marshal(srvs)

	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
