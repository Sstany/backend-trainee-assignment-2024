package router

import (
	"banney/app/db"

	"github.com/muesli/cache2go"
	"go.uber.org/zap"
)

type Router struct {
	DB          *db.Client
	Logger      *zap.Logger
	BannerCache *cache2go.CacheTable
}

func New(dbClient *db.Client, logger *zap.Logger) *Router {
	return &Router{
		DB:          dbClient,
		Logger:      logger,
		BannerCache: cache2go.Cache("BannerCache"),
	}
}
