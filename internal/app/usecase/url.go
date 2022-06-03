package usecase

import (
	"context"
	"fmt"

	"github.com/e4t4g/URLback/internal/app/repository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type URLData struct {
	ID       int
	FullURL  string
	ShortURL string
	Counter  int64
}

type UseCase interface {
	Create(ctx context.Context, fullURL *URLData) (*URLData, error)
	Redirect(ctx context.Context, token string) (*URLData, error)
	GetStat(ctx context.Context, id int) (*URLData, error)
}

type useCaseURL struct {
	repository repository.Repository
	logger     *zap.SugaredLogger
}

func New(repo repository.Repository, logger *zap.SugaredLogger) *useCaseURL {
	return &useCaseURL{repository: repo, logger: logger}
}

func (usc *useCaseURL) Create(ctx context.Context, fullURL *URLData) (*URLData, error) {
	fullURL.ShortURL = tokenGen()

	result, err := usc.repository.Create(ctx, (*repository.URLData)(fullURL))
	if err != nil {
		usc.logger.Error(err)
		return nil, fmt.Errorf("usecase error: can not create: %s", err)
	}

	return (*URLData)(result), nil
}

func (usc *useCaseURL) Redirect(ctx context.Context, token string) (*URLData, error) {
	redirect, err := usc.repository.FindByToken(ctx, token)
	if err != nil {
		usc.logger.Error(err)
		return nil, fmt.Errorf("usecase error: can not redirect: %s", err)
	}

	count := redirect.Counter + 1

	err = usc.repository.UpdateCounter(ctx, count, redirect.ShortURL)
	if err != nil {
		usc.logger.Error(err)
		return nil, fmt.Errorf("usecase error: can not update counter: %s", err)
	}

	return (*URLData)(redirect), nil
}

func (usc *useCaseURL) GetStat(ctx context.Context, id int) (*URLData, error) {
	getStat, err := usc.repository.FindByID(ctx, id)
	if err != nil {
		usc.logger.Error(err)
		return nil, fmt.Errorf("usecase error: can not get stat: %s", err)
	}

	return (*URLData)(getStat), nil
}

func tokenGen() string {
	id := uuid.New()
	return id.String()[:8]
}
