package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	secretKey = []byte("secret")
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(email string) (string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
			 jwt.MapClaims{
				"email": email,
				"exp": time.Now().Add(time.Hour * 24).Unix(),
			 })
	
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return ""
	}

	return tokenString

			}

func ValidateToken(tokenString string) (bool) {
	jwtToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return false
	}

	return jwtToken.Valid
}