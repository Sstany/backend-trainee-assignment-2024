package banner

import (
	"banney/sdk"
	"banney/sdk/authorize"
	"banney/sdk/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (r *BannerRouter) create(ctx *gin.Context) {
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

	var banner models.Banner
	err = json.NewDecoder(ctx.Request.Body).Decode(&banner)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &models.ServerError{Error: err.Error()})
		return
	}

	r.Logger.Debug("banner info", zap.Any("body", banner))

	bannerID, err := r.DB.CreateBanner(ctx, &banner)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, &models.ServerError{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, &models.BannerCreated{ID: bannerID})

}
