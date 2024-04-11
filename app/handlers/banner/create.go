package banner

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *BannerRouter) create(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ok")
}
