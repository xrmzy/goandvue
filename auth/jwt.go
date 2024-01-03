package auth

import "github.com/golang-jwt/jwt/v5"

type JWTService interface {
	GenerateToken(userID string) (string, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("RMZSTARTUP_S3crEt_KeYs")

func (s *jwtService) GenerateToken(userID string) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func NewJWTService() *jwtService {
	return &jwtService{}
}
