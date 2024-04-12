package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"banney/sdk"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type Client struct {
	cli    *sql.DB
	logger *zap.Logger
}

func NewClient(log *zap.Logger) *Client {
	db, err := sql.Open("postgres", os.Getenv(sdk.EnvPostgres))
	if err != nil {
		panic(err)
	}

	return &Client{
		cli:    db,
		logger: log,
	}
}

func (r *Client) Start() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	var err error

	_, err = r.cli.ExecContext(ctx, queryCreateTagTable)
	if err != nil && !sdk.IsDublicateTableErr(err) {
		r.logger.Error("creating tag table", zap.Error(err))

		return fmt.Errorf("creating tag table: %w", err)
	}

	_, err = r.cli.ExecContext(ctx, queryCreateFeatureTable)
	if err != nil && !sdk.IsDublicateTableErr(err) {
		r.logger.Error("creating feature table", zap.Error(err))

		return fmt.Errorf("creating feature table: %w", err)
	}

	_, err = r.cli.ExecContext(ctx, queryCreateBannerTable)
	if err != nil && !sdk.IsDublicateTableErr(err) {
		r.logger.Error("creating banner table", zap.Error(err))

		return fmt.Errorf("creating banner table: %w", err)
	}

	_, err = r.cli.ExecContext(ctx, queryCreateBannerTagTable)
	if err != nil && !sdk.IsDublicateTableErr(err) {
		r.logger.Error("creating banner_tag table", zap.Error(err))

		return fmt.Errorf("creating banner_tag table: %w", err)
	}

	return nil
}
