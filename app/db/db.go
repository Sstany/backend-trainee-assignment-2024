package db

import (
	"database/sql"
	"os"

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

	return &Client{cli: db, logger: log}
}
