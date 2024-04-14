package authorize

import (
	"errors"

	"banney/sdk"
	"banney/sdk/models"

	"github.com/golang-jwt/jwt"
)

var (
	errTokenInvalid    = errors.New("token invalid")
	errClaimsWrongType = errors.New("claims is wrong typed")
)

func AuthToken(token string) (*models.Claims, error) {
	tokenJWT, err := jwt.ParseWithClaims(token, &models.Claims{}, Check)
	if err != nil {
		return nil, err
	}

	if !tokenJWT.Valid {
		return nil, errTokenInvalid
	}

	claims, ok := tokenJWT.Claims.(*models.Claims)
	if !ok {
		return nil, errClaimsWrongType
	}

	return claims, nil
}

func Check(token *jwt.Token) (interface{}, error) {
	return sdk.Secret, nil
}
