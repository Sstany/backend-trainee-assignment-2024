package banner

import (
	"banney/sdk/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (r *BannerRouter) delete(ctx *gin.Context) {
	var banner models.Banner

	err := json.NewDecoder(ctx.Request.Body).Decode(&banner)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &models.ServerError{Error: err.Error()})
		return
	}

	r.Logger.Debug("banner info", zap.Any("body", banner))

	bannerID, err := r.DB.DeleteBanner(ctx, &banner)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, &models.ServerError{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, &models.BannerDeleted{ID: bannerID})

}
