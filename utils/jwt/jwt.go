package jwt

import (
	"time"

	"github.com/agusheryanto182/go-raide-hailing/config"
	"github.com/agusheryanto182/go-raide-hailing/utils/customErr"
	"github.com/golang-jwt/jwt/v5"
)

type JWTInterface interface {
	GenerateJWT(ID, username, role string) (string, error)
	ValidateToken(tokenString string) (*JWTPayload, error)
}

type JWTService struct {
	cfg config.Global
}

func NewJWTService(cfg config.Global) JWTInterface {
	return &JWTService{
		cfg: cfg,
	}
}

type JWTPayload struct {
	Id       string
	Username string
	Role     string
}

type JWTClaims struct {
	Id       string
	Username string
	Role     string
	jwt.RegisteredClaims
}

func (s *JWTService) GenerateJWT(id, username, role string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		Id:       id,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
		},
	})

	tokenString, err := token.SignedString([]byte(s.cfg.JwtSecretKey))
	return tokenString, err
}

func (s *JWTService) ValidateToken(tokenString string) (*JWTPayload, error) {
	claims := &JWTClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JwtSecretKey), nil
	})
	if err != nil {
		return nil, customErr.NewUnauthorizedError("Access denied: " + err.Error())
	}

	if claims.RegisteredClaims.ExpiresAt.Before(time.Now()) {
		return nil, customErr.NewUnauthorizedError("Access denied: token expired")
	}

	payload := &JWTPayload{
		Id:       claims.Id,
		Username: claims.Username,
		Role:     claims.Role,
	}

	return payload, nil
}
