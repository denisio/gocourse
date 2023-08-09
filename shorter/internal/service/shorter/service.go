package shorter

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"shorter/internal/domain"
	"shorter/internal/provider"

	"github.com/gorilla/mux"
)

type Service struct {
	storage provider.StorageProvider
	ctx     context.Context
}

func NewService(storage provider.StorageProvider) *Service {
	return &Service{storage: storage}
}

func (s *Service) Start(ctx context.Context) error {
	s.ctx = ctx
	return nil
}

func (s *Service) Stop() {
}

func isValidUrl(token string) bool {
	_, err := url.ParseRequestURI(token)

	if err != nil {
		log.Printf("isValidUrl: %v", err)
		return false
	}

	u, err := url.Parse(token)

	if err != nil || u.Host == "" {
		log.Printf("isValidUrl: %v", err)
		return false
	}

	return true
}

func (s *Service) IndexPage(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles(domain.AppDir + "/templates/index.html")
	result := domain.Result{}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == "POST" {
		if !isValidUrl(r.FormValue("url")) {
			result.Status = "Ссылка имеет неправильный формат!"
			result.Url = ""
		} else {
			result.Url = r.FormValue("url")
			key, err := s.storage.Create(s.ctx, result.Url)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			result.Key = key
		}
	}

	templ.Execute(w, result)
}

func (s *Service) RedirectTo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	url, err := s.storage.GetByID(s.ctx, vars["key"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//fmt.Fprintf(w, "<script>location='%s';</script>", url)
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
