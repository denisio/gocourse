package main

import (
	"shorter/internal/app"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	application := app.NewApp()
	application.Run()
}
