package middleware

import (
	"fmt"

	api_structure "github.com/FurkanSamaraz/Golang-Chat/internal/pkg/structures"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type JWTMiddleware struct {
	SecretKey string
}

func NewJWTMiddleware(secretKey string) *JWTMiddleware {
	return &JWTMiddleware{
		SecretKey: secretKey,
	}
}

func (j *JWTMiddleware) Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Eksik yetkilendirme belirteci"})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("beklenmeyen imzalama yöntemi: %v", token.Header["alg"])
			}
			return []byte(j.SecretKey), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(api_structure.AuthErrorResponse{
				Message: "Geçersiz yetkilendirme belirteci",
			})
		}

		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(api_structure.AuthErrorResponse{
				Message: "Geçersiz yetkilendirme belirteci",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(api_structure.AuthErrorResponse{
				Message: "Geçersiz yetkilendirme bilgileri",
			})
		}

		user := &api_structure.Claims{
			ID:   claims["id"].(float64),
			Name: claims["name"].(string),
		}

		c.Locals("user", user)

		return c.Next()
	}
}
