package auth_service

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ClaimValues struct {
	ID string
}

type Claims struct {
	jwt.RegisteredClaims
}

type IAuthService interface {
	GenerateJWT(data ClaimValues) (string, error)
	VerifyJWT(token string) (Claims, error)
}

type AuthService struct{}

const jwtExpiresHours = 48 * time.Hour

func (s *AuthService) GenerateJWT(claims ClaimValues) (string, error) {
	secret, exists := os.LookupEnv("JWT_SECRET")
	if !exists {
		return "", errors.New("cannot authenticate, missing ENV dependencies")
	}

	c := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        claims.ID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtExpiresHours)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, c)
	return token.SignedString(secret)
}

func (s *AuthService) VerifyJWT(token string) (Claims, error) {
	secret, exists := os.LookupEnv("JWT_SECRET")
	if !exists {
		return Claims{}, errors.New("cannot authorize, missing ENV dependencies")
	}

	result, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return secret, nil
	})
	if err != nil {
		return Claims{}, err
	}
	claims, ok := result.Claims.(Claims)
	if !ok {
		return Claims{}, errors.New("cannot verify JWT. Failed to parse claims")
	}

	return claims, nil
}

func New() *AuthService {
	return &AuthService{}
}
