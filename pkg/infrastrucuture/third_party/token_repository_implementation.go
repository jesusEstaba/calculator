package third_party

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jesusEstaba/calculator/internal"
	"github.com/jesusEstaba/calculator/pkg/domain"
)

type TokenRepositoryImplementation struct{}

func NewTokenRepositoryImplementation() domain.TokenRepository {
	return &TokenRepositoryImplementation{}
}

func (r *TokenRepositoryImplementation) Generate(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(internal.Config.JWTSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (r *TokenRepositoryImplementation) Verify(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token: parse")
		}

		return []byte(internal.Config.JWTSecret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userID, ok := claims["user_id"].(string); ok {
			return userID, nil
		}
	}

	return "", fmt.Errorf("invalid token: bad mapping")
}
