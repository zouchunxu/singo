package jwt

import (
	j "github.com/dgrijalva/jwt-go"
	"singo/config"
	"time"
)

func GenerateToken(id int) (string, error) {
	cfg := config.GetConfig()

	now := time.Now()

	seconds := cfg.GetInt64("jwt.expire")
	expire := now.Add(time.Duration(seconds * int64(time.Second)))

	claims := Claims{
		id,
		j.StandardClaims{
			ExpiresAt: expire.Unix(),
			Issuer:    "go",
		},
	}

	tokenClaims := j.NewWithClaims(j.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(cfg.GetString("jwt.secret")))

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	cfg := config.GetConfig()

	tokenClaims, err := j.ParseWithClaims(token, &Claims{}, func(token *j.Token) (interface{}, error) {
		return []byte(cfg.GetString("jwt.secret")), nil
	})

	if err != nil {
		if ve, ok := err.(*j.ValidationError); ok {
			if ve.Errors&j.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&j.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&j.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}

	return nil, err
}
