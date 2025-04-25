package jwt

import (
	"time"

	"github.com/g0shi4ek/store/config"
	"github.com/g0shi4ek/store/internal/store/domain"
	"github.com/golang-jwt/jwt"
)

func CreateNewToken(user *domain.User, cfg *config.Config) (string, error) {
	key := []byte(cfg.StoreConf.SecretKey)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Username,                         // Subject (user identifier)
		"iss": "app",                                 // Issuer
		"aud": user.Role,                             // Audience (user role)
		"exp": time.Now().Add(time.Hour * 72).Unix(), // Expiration time
		"iat": time.Now().Unix(),                     // Issued at
	})

	tokenString, err := claims.SignedString(key)
	if err != nil {
		return "", err
	}
	return "Bearer " + tokenString, nil
}


func VerifyToken(tokenString string, cfg *config.Config) (*jwt.Token, error) {

	claims := jwt.StandardClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.StoreConf.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return token, nil
}
