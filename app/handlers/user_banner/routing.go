package userbanner

import (
	"banney/app/router"

	"github.com/gin-gonic/gin"
)

type UserBannerRouter struct {
	*router.Router
}

func NewBannerRouter(pr *router.Router) *UserBannerRouter {
	b := &UserBannerRouter{
		pr,
	}
	b.Logger = b.Logger.Named("banner")

	return b
}

func AttachToGroup(pr *router.Router, parentGroup *gin.RouterGroup) {
	userBannerRouter := NewBannerRouter(pr)

	parentGroup.GET("", userBannerRouter.get)
}
