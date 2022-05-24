package util

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SecretKey = "secret"

func GenerateJwt(Issuer string, Id uint) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    Issuer,
		Subject:   strconv.Itoa(int(Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})
	return claims.SignedString([]byte(SecretKey))
}
func ParseJwt(cookie string) (uint, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		return 0, err
	}
	claims := token.Claims.(*jwt.StandardClaims)
	id, _ := strconv.Atoi(claims.Subject)

	return uint(id), nil
}
