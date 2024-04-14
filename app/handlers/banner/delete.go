package banner

import (
	"banney/app/db"
	"banney/sdk"
	"banney/sdk/authorize"
	"banney/sdk/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (r *BannerRouter) delete(ctx *gin.Context) {

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

	bannerID := ctx.Param("id")

	err = r.DB.DeleteBanner(ctx, bannerID)

	if err != nil {
		if errors.Is(err, db.ErrBannerNotExists) {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, &models.ServerError{Error: err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
