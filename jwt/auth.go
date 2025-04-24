package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("H^$$^NZ^H!D")

type Claims struct {
	Email  string `json:"email"`
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(email string, userId string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Second)

	claims := &Claims{
		Email:  email,
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWTToken(tokenStr string) bool {
	if tokenStr == "" {
		return false
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return false
	}

	if err := claims.Valid(); err != nil {
		return false
	}

	return true
}

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Get("Authorization")
		if tokenStr == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
			tokenStr = tokenStr[7:]
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		if !token.Valid {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		c.Locals("email", claims.Email)
		// c.Locals("role", claims.Role)
		c.Locals("user_id", claims.UserId)
		return c.Next()
	}
}

func JWTMiddlewareValidation(token string) bool {
	if token == "" {
		return false
	}

	claims := &Claims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !parsedToken.Valid {
		return false
	}

	return true
}
