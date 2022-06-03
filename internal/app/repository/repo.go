package repository

import "context"

type URLData struct {
	ID       int    `db:"id"`
	FullURL  string `db:"full_url"`
	ShortURL string `db:"short_url"`
	Counter  int64  `db:"counter"`
}

type Repository interface {
	Create(ctx context.Context, url *URLData) (*URLData, error)
	FindByToken(ctx context.Context, token string) (*URLData, error)
	UpdateCounter(ctx context.Context, counter int64, shortURL string) error
	FindByID(ctx context.Context, id int) (*URLData, error)
}
