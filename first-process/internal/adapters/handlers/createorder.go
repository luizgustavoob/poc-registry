package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/luizgustavoob/first-process/internal/core/entities"
	"github.com/luizgustavoob/first-process/internal/core/ports"
)

type createOrderHandler struct {
	o ports.Order
}

func NewCreateOrderHandler(o ports.Order) *createOrderHandler {
	return &createOrderHandler{
		o: o,
	}
}

func (h *createOrderHandler) Method() string {
	return http.MethodPost
}

func (h *createOrderHandler) Pattern() string {
	return "/commands"
}

func (h *createOrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var order entities.Create

	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		result := make(map[string]string)
		result["error"] = err.Error()
		json, _ := json.Marshal(result)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
		return
	}

	h.o.Create()

	body, _ := io.ReadAll(r.Body)
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}
