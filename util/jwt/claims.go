package jwt

import j "github.com/dgrijalva/jwt-go"

type Claims struct {
	ID int `json:"id"`
	j.StandardClaims
}
