package auth

import (
	"banney/sdk"
	"banney/sdk/models"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (r *AuthRouter) login(ctx *gin.Context) {
	var loginForm models.LoginForm

	err := json.NewDecoder(ctx.Request.Body).Decode(&loginForm)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &models.ServerError{Error: err.Error()})
		return
	}

	user, err := r.DB.GetUser(ctx, &loginForm)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, &models.ServerError{Error: err.Error()})
		return
	}

	claims := models.Claims{
		UserID:  user.ID,
		IsAdmin: user.IsAdmin,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	jwtToken, err := token.SignedString(sdk.Secret)

	ctx.JSON(http.StatusOK, &models.Access{Token: jwtToken})
}
