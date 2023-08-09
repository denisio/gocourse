package service

import (
	"context"
	"net/http"
)

type ShorterService interface {
	Start(ctx context.Context) error
	Stop()

	IndexPage(w http.ResponseWriter, r *http.Request)
	RedirectTo(w http.ResponseWriter, r *http.Request)
}
