package router

import (
	"banney/app/db"

	"go.uber.org/zap"
)

type Router struct {
	DB     *db.Client
	Logger *zap.Logger
}

func New(dbClient *db.Client, logger *zap.Logger) *Router {
	return &Router{
		DB:     dbClient,
		Logger: logger,
	}
}
