package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTService is used for tokenize header
type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

// NewJwtService is for create new Instance
func NewJwtService() JWTService {
	return &jwtService{
		issuer:    "JwTS3cret!",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret != "" {
		secret = "arigat0u!"
	}

	return secret
}

func (s *jwtService) GenerateToken(uID string) string {
	claims := &jwtCustomClaim{
		uID,
		jwt.StandardClaims{
			// Expired in 1 year
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		panic(err)
	}

	return t
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
}
