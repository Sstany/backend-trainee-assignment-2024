package banner

import (
	"banney/sdk/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (r *BannerRouter) update(ctx *gin.Context) {
	var banner models.Banner

	bannerID := ctx.Param("id")
	err := json.NewDecoder(ctx.Request.Body).Decode(&banner)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &models.ServerError{Error: err.Error()})
		return
	}
	banner.ID, err = strconv.Atoi(bannerID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &models.ServerError{Error: err.Error()})
		return
	}

	r.Logger.Debug("banner info", zap.Any("body", banner))

	banner.ID, err = r.DB.UpdateBanner(ctx, &banner)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, &models.ServerError{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, &models.BannerCreated{ID: banner.ID})

}
