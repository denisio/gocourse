package handlers

import "net/http"

type Handlers interface {
	IndexPage(w http.ResponseWriter, r *http.Request)
	RedirectTo(w http.ResponseWriter, r *http.Request)
}
