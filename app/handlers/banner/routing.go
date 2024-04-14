package banner

import (
	"banney/app/router"

	"github.com/gin-gonic/gin"
)

type BannerRouter struct {
	*router.Router
}

func NewBannerRouter(pr *router.Router) *BannerRouter {
	b := &BannerRouter{
		pr,
	}
	b.Logger = b.Logger.Named("banner")

	return b
}

func AttachToGroup(pr *router.Router, parentGroup *gin.RouterGroup) {
	bannerRouter := NewBannerRouter(pr)

	parentGroup.POST("", bannerRouter.create)
	// parentGroup.GET("/")
	// parentGroup.PATCH("/:id")
	parentGroup.DELETE("/:id", bannerRouter.delete)
}
