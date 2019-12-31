package jwt

import "github.com/dgrijalva/jwt-go"

func TokenGenerate(mapClaims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	return token.SignedString([]byte("secret"))
}
