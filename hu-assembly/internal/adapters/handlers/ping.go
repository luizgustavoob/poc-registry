package handlers

import (
	"net/http"
)

type pingHandler struct{}

func NewPingHandler() *pingHandler {
	return &pingHandler{}
}

func (h *pingHandler) Method() string {
	return http.MethodGet
}

func (h *pingHandler) Pattern() string {
	return "/ping"
}

func (h *pingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
