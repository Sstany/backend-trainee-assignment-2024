package banner

import (
	"banney/sdk"
	"banney/sdk/authorize"
	"banney/sdk/models"
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (r *BannerRouter) list(ctx *gin.Context) {
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

	// No admin rights to perform request
	if !claims.IsAdmin {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	tagIDStr := ctx.Query(sdk.QueryTagID)
	featureIDStr := ctx.Query(sdk.QueryFeatureID)
	offsetStr := ctx.Query(sdk.QueryOffset)
	limitStr := ctx.Query(sdk.QueryLimit)

	// TagID conversion
	var tagID int
	if tagIDStr != "" {
		tagID, err = strconv.Atoi(tagIDStr)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}

	// TagID conversion
	var featureID int
	if featureIDStr != "" {
		featureID, err = strconv.Atoi(featureIDStr)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}

	// Offset conversion
	var offset int
	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}

	// Limit conversion
	limit := sdk.DefaultLimit
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}

	// Offset and Limit validation
	if limit > sdk.MaxLimit || offset > sdk.MaxLimit {
		r.Logger.Error(
			"max limit exceeded",
			zap.Int("limit", limit),
			zap.Int("offset", offset),
		)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if limit < 0 || offset < 0 {
		r.Logger.Error(
			"less than zero",
			zap.Int("limit", limit),
			zap.Int("offset", offset),
		)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if tagIDStr != "" && featureIDStr != "" {
		banner, err := r.DB.GetBanner(ctx, featureIDStr, tagIDStr)
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

		ctx.JSON(http.StatusOK, &[]models.Banner{*banner})
		return
	}

	filter := &models.BannerFilter{
		FeatureID: featureID,
		TagID:     tagID,
		Offset:    offset,
		Limit:     limit,
	}

	banners, err := r.DB.FilterBanners(ctx, filter)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, &models.ServerError{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &banners)
}
