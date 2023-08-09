package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
)

type Provider struct {
	db *sql.DB
}

func NewProvider(db *sql.DB) *Provider {
	return &Provider{db: db}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func shorting() string {
	b := make([]byte, 5)

	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}

func (p *Provider) Create(ctx context.Context, url string) (string, error) {
	id := shorting()

	if _, err := p.db.Exec("insert into links (link, short) values ($1, $2)", url, id); err != nil {
		return "", fmt.Errorf("insert failed: %v", err)
	}

	return id, nil
}

func (p *Provider) GetByID(ctx context.Context, id string) (string, error) {
	var url string

	rows := p.db.QueryRow("select link from links where short=$1 limit 1", id)
	err := rows.Scan(&url)

	if err != nil {
		return "", fmt.Errorf("url not found: %s", id)
	}

	return url, nil
}

func (p *Provider) Close() {
	p.db.Close()
}
