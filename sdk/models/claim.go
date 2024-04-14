package models

import "github.com/golang-jwt/jwt"

type Claims struct {
	jwt.StandardClaims `json:",inline"`

	UserID  int  `json:"user_id"`
	IsAdmin bool `json:"is_admin"`
}
