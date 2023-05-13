package structures

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	ID   float64 `json:"id"`
	Name string  `json:"name"`
	jwt.StandardClaims
}

type AuthErrorResponse struct {
	Message string `json:"message"`
}
