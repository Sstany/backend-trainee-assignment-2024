package banner

import (
	"banney/app/db"
	"banney/sdk/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *BannerRouter) delete(ctx *gin.Context) {

	bannerID := ctx.Param("id")

	err := r.DB.DeleteBanner(ctx, bannerID)

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
