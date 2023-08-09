package shorter

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"shorter/internal/service/mocks"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestIndexPage(t *testing.T) {
	// Set up a mock storage provider
	storProv := new(mocks.ProviderMock)
	storProv.On("Create", mock.Anything, mock.Anything).Return("12345", nil)

	// Create a new Service instance
	s := &Service{storage: storProv}
	s.Start(context.Background())

	// Create a mock HTTP request with a valid URL
	data := url.Values{}
	data.Set("url", "https://ya.ru")

	body := strings.NewReader(data.Encode())
	req := httptest.NewRequest("POST", "/", body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Create a mock HTTP response
	w := httptest.NewRecorder()

	// Call the IndexPage function
	s.IndexPage(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	if !strings.Contains(w.Body.String(), "12345") {
		t.Errorf("Expected body %s, but got %s", "12345", w.Body.String())
	}

	// Check that the Create method of the storage provider was called
	storProv.AssertCalled(t, "Create", mock.Anything, mock.Anything)
}

func TestRedirect(t *testing.T) {
	// Set up a mock storage provider
	storProv := new(mocks.ProviderMock)
	storProv.On("GetByID", mock.Anything, mock.Anything).Return("http://ya.ru", nil)

	// Create a new Service instance
	s := &Service{storage: storProv}
	s.Start(context.Background())

	// Create a mock HTTP request
	req := httptest.NewRequest("GET", "/12345", nil)

	// Create a mock HTTP response
	w := httptest.NewRecorder()

	// Call the Redirect function
	s.RedirectTo(w, req)

	// Check the response status code
	if w.Code != http.StatusMovedPermanently {
		t.Errorf("Expected status code %d, but got %d", http.StatusMovedPermanently, w.Code)
	}

	if w.Header().Get("Location") != "http://ya.ru" {
		t.Errorf("Expected Location header %s, but got %s", "http://ya.ru", w.Header().Get("Location"))
	}

	// Check that the Create method of the storage provider was called
	storProv.AssertCalled(t, "GetByID", mock.Anything, mock.Anything)
}
