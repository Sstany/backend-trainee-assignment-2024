package auth

import (
	"banney/app/router"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	*router.Router
}

func NewAuthRouter(pr *router.Router) *AuthRouter {
	b := &AuthRouter{
		pr,
	}
	b.Logger = b.Logger.Named("auth")
	return b
}
func AttachToGroup(pr *router.Router, parentGroup *gin.RouterGroup) {
	authRouter := NewAuthRouter(pr)
	parentGroup.POST("/register", authRouter.register)
	parentGroup.POST("/login", authRouter.login)
}
