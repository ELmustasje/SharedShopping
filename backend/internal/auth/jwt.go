package auth

import (
	"os"
	"thomas-barth/SharedShopping/db/READ"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = os.Getenv("JWT_SECRET")

type Claims struct {
	UserID  uint   `json:"user_id"`
	ListIDs []uint `json:"listIDs"`
	jwt.RegisteredClaims
}

// Create JWT token
func GenerateToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	lists, err := READ.ReadUserLists(userID)
	if err != nil {
		return "", err // or handle the error appropriately
	}

	listIDs := make([]uint, 0, len(lists))
	for _, list := range lists {
		listIDs = append(listIDs, list.ID)
	}

	claims := &Claims{
		UserID:  userID,
		ListIDs: listIDs,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Parse JWT token
func ParseToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
