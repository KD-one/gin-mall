package util

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// Claims 用户信息
type Claims struct {
	UserId uint
	jwt.RegisteredClaims
}

// key 用于签名的密钥
var key = []byte("your-secret-key")

// ReleaseToken 生成token
func ReleaseToken(userId uint) (string, error) {
	claim := &Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "127.0.0.1",
			Subject:   "token",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	return token, claims, err
}
