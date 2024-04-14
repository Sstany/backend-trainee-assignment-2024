package router

import (
	"banney/sdk"
	"banney/sdk/models"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (r *Router) CacheBanner(banner *models.Banner, userID int) {
	logger := r.Logger.Named("cache").With(
		zap.String("traceID", uuid.NewString()),
		zap.Int("bannerID", banner.ID),
		zap.Int("userID", userID),
	)

	logger.Info("starting banner caching")

	for _, tagID := range banner.TagIDs {
		key := sdk.CreateKey(banner.FeatureID, tagID)
		r.BannerCache.Add(key, sdk.DefaulCacheLifetime, banner)
	}

	logger.Info("finished banner caching")
}
