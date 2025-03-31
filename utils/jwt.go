package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jpeccia/bebi-delivery-server/internal/models"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Role string `json:"role"`
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(user *models.User) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: fmt.Sprintf("%d", user.ID),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer: "bebi-delivery-api",
		},
		Role: string(user.Role),
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (uint, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return 0, "", err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return 0, "", fmt.Errorf("Não foi possivel converter os claims")
	}

	id, err := strconv.ParseUint(claims.Subject, 10, 64)
	if err != nil {
		return 0, "", err
	}

	return uint(id), claims.Role, nil
}
