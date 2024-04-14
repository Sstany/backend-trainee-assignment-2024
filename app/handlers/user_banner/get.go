package userbanner

import (
	"database/sql"
	"errors"
	"net/http"

	"banney/sdk"
	"banney/sdk/authorize"
	"banney/sdk/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (r *UserBannerRouter) get(ctx *gin.Context) {
	token := ctx.GetHeader(sdk.HeaderToken)
	claims, err := authorize.AuthToken(token)
	if err != nil {
		r.Logger.Error("authorize", zap.Error(err))
	}

	// No or malformed token
	if claims == nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tagID := ctx.Query(sdk.QueryTagID)
	featureID := ctx.Query(sdk.QueryFeatureID)
	lastRevisionStr := ctx.Query(sdk.QueryUseLastRevision)

	var lastRevision bool

	if lastRevisionStr == "true" {
		lastRevision = true
	}

	var banner *models.Banner

	cacheKey := sdk.CreateKeyFromString(featureID, tagID)

	res, err := r.BannerCache.Value(cacheKey)
	if err != nil || lastRevision {
		r.Logger.Info(
			"banner not found in cache",
			zap.String("featureID", featureID),
			zap.String("tagID", tagID),
		)

		banner, err = r.DB.GetBanner(ctx, featureID, tagID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				ctx.AbortWithStatus(http.StatusNotFound)
				return
			}

			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				&models.ServerError{Error: err.Error()})
			return
		}

		// Perform caching in other goruitine to prevent response delay
		go r.CacheBanner(banner, claims.UserID)
	} else {
		var ok bool
		banner, ok = res.Data().(*models.Banner)
		if !ok {
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				&models.ServerError{Error: "Conversion of cached element failed"})
			return
		}
	}

	if !banner.IsActive && !claims.IsAdmin {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	ctx.JSON(http.StatusOK, &banner.Content)
}
