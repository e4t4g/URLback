package sqlite

import (
	"context"
	"fmt"

	"github.com/e4t4g/URLback/internal/app/repository"

	"github.com/jmoiron/sqlx"
	_ "github.com/mxk/go-sqlite/sqlite3"
	"go.uber.org/zap"
)

type repoURL struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

const querySelect = "SELECT * FROM url WHERE "

func New(db *sqlx.DB, logger *zap.SugaredLogger) repository.Repository {
	return &repoURL{
		db:     db,
		logger: logger,
	}
}

func (rdb *repoURL) Create(ctx context.Context, url *repository.URLData) (*repository.URLData, error) {
	var URLData repository.URLData

	query := "INSERT INTO url (full_url, short_url, counter) VALUES (?, ?, ?)"
	statement, err := rdb.db.Prepare(query)
	if err != nil {
		rdb.logger.Error(err)
		return nil, fmt.Errorf("failed to create: %w", err)
	}
	statement.QueryRow(url.FullURL, url.ShortURL, 0)

	query = querySelect + "short_url = ?"

	err = rdb.db.GetContext(ctx, &URLData, query, url.ShortURL)
	if err != nil {
		rdb.logger.Error(err)
		return nil, fmt.Errorf("failed to create: %w", err)
	}

	return &URLData, nil
}

func (rdb *repoURL) FindByToken(ctx context.Context, token string) (*repository.URLData, error) {
	var URLData repository.URLData

	statQuery := querySelect + "short_url = ?"

	err := rdb.db.GetContext(ctx, &URLData, statQuery, token)
	if err != nil {
		rdb.logger.Error(err)
		return nil, fmt.Errorf("failed to find by token: %w", err)
	}

	return &URLData, nil
}

func (rdb *repoURL) FindByID(ctx context.Context, id int) (*repository.URLData, error) {
	var URLData repository.URLData

	statQuery := querySelect + "id = ?"

	err := rdb.db.GetContext(ctx, &URLData, statQuery, id)
	if err != nil {
		rdb.logger.Error(err)
		return nil, fmt.Errorf("failed to find by id: %w", err)
	}

	return &URLData, nil
}

func (rdb *repoURL) UpdateCounter(ctx context.Context, counter int64, shortURL string) error {
	query := "UPDATE url SET counter = ? WHERE short_url = ? "
	statement, err := rdb.db.Prepare(query)
	if err != nil {
		rdb.logger.Error(err)
		return fmt.Errorf("failed to update counter: %w", err)
	}
	statement.QueryRow(counter, shortURL)
	return nil
}
