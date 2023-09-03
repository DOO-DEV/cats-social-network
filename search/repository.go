package search

import (
	"context"
	"meower/schema"
)

type Repository interface {
	Close()
	InsertMeow(ctx context.Context, meow schema.Meow) error
	ListMeows(ctx context.Context, query string, skip, take uint64) ([]schema.Meow, error)
}

var implement Repository

func SetRepository(repo Repository) {
	implement = repo
}

func Close() {
	implement.Close()
}

func InsertMeow(ctx context.Context, meow schema.Meow) error {
	return implement.InsertMeow(ctx, meow)
}

func ListMeows(ctx context.Context, query string, skip, take uint64) ([]schema.Meow, error) {
	return implement.ListMeows(ctx, query, skip, take)
}
