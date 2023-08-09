package app

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"shorter/internal/handlers"
	"shorter/internal/provider"
	"shorter/internal/provider/sqlite"
	"shorter/internal/service"
	"shorter/internal/service/shorter"
	"syscall"

	"github.com/gorilla/mux"
)

type App struct {
	storage provider.StorageProvider
	service service.ShorterService
	router  *mux.Router
}

func NewApp() *App {
	db, err := sql.Open("sqlite3", "shorter.db")

	if err != nil {
		log.Fatal(err)
	}

	var storage provider.StorageProvider = sqlite.NewProvider(db)
	var service service.ShorterService = shorter.NewService(storage)

	router := mux.NewRouter()

	if err = handlers.Register(router, service); err != nil {
		log.Fatal(err)
	}

	return &App{storage: storage, service: service, router: router}
}

func (a *App) Run() {
	ctx, cancel := context.WithCancel(context.Background())

	err := a.service.Start(ctx)

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Fatal(http.ListenAndServe(":8181", a.router))
	}()

	log.Print("Service started")

	a.waitGracefulShutdown(cancel)
}

func (a *App) waitGracefulShutdown(cancel context.CancelFunc) {
	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGTERM, os.Interrupt,
	)

	log.Printf("Caught signal %s. Shutting down...", <-quit)

	cancel()

	a.service.Stop()
	a.storage.Close()
}
