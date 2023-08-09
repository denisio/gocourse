package restapi

import (
	"net/http"
	"shorter/internal/service"
)

type Handlers struct {
	service service.ShorterService
}

func NewHandlers(service service.ShorterService) *Handlers {
	return &Handlers{service: service}
}

func (h *Handlers) IndexPage(w http.ResponseWriter, r *http.Request) {
	h.service.IndexPage(w, r)
}

func (h *Handlers) RedirectTo(w http.ResponseWriter, r *http.Request) {
	h.service.RedirectTo(w, r)
}
