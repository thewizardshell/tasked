package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenManager struct {
	secret    string
	expiryHrs int
}

func NewTokenManager(secret string, expiryHrs int) *TokenManager {
	return &TokenManager{
		secret:    secret,
		expiryHrs: expiryHrs,
	}
}

type CustomClaims struct {
	UserID   int64  `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (tm *TokenManager) GenerateToken(userID int64, email, username string) (string, error) {
	claims := CustomClaims{
		UserID:   userID,
		Email:    email,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(tm.expiryHrs))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "tasked-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(tm.secret))
}

func (tm *TokenManager) ValidateToken(tokenString string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(tm.secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func (tm *TokenManager) RefreshToken(oldToken string) (string, error) {
	claims, err := tm.ValidateToken(oldToken)
	if err != nil {
		return "", err
	}

	return tm.GenerateToken(claims.UserID, claims.Email, claims.Username)
}
