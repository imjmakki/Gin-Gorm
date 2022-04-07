package service

type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
}
