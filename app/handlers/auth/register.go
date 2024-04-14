package auth

import (
	"encoding/json"
	"net/http"

	"banney/sdk"
	"banney/sdk/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (r *AuthRouter) register(ctx *gin.Context) {
	var user models.RegisterForm
	err := json.NewDecoder(ctx.Request.Body).Decode(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &models.ServerError{Error: err.Error()})
		return
	}
	userID, err := r.DB.CreateUser(ctx, &user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, &models.ServerError{Error: err.Error()})
		return
	}

	claims := models.Claims{
		UserID:  userID,
		IsAdmin: user.IsAdmin,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	jwtToken, err := token.SignedString(sdk.Secret)

	ctx.JSON(http.StatusOK, &models.Access{Token: jwtToken})
}
