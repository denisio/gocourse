package handlers

import (
	"shorter/internal/handlers/restapi"
	"shorter/internal/service"

	"github.com/gorilla/mux"
)

const (
	IndexPagePath  = "/"
	RedirectToPath = "/{key}"
)

func Register(router *mux.Router, service service.ShorterService) error {
	h := restapi.NewHandlers(service)

	router.HandleFunc(IndexPagePath, h.IndexPage)
	router.HandleFunc(RedirectToPath, h.RedirectTo)

	return nil
}
