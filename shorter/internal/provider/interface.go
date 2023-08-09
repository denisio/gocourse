package provider

import "context"

type StorageProvider interface {
	Create(ctx context.Context, url string) (string, error)
	GetByID(ctx context.Context, id string) (string, error)
	Close()
}
