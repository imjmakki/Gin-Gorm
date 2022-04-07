package service

import "os"

type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "MJ",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "mj"
	}
	return secretKey
}

func (s *jwtService) GenerateToken(UserID string) string {

}
