package util

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// 利用邮箱验证身份时需要将邮件中返回的链接进行加密
// 防止篡改：未经加密的链接可能被恶意用户拦截并篡改其中的内容，例如更改待验证用户的ID或验证状态。通过加密Token，即使链接内容被截获，也不能轻易地理解或修改其原始信息。
// 防止重放攻击：加密Token通常具有有效期和一次性使用的特性，这样即使攻击者获取了Token，并试图重复使用它来验证身份，也会因为Token已失效或已被消耗而无法成功。

type EmailMsg struct {
	UserId        uint   `json:"user_id" form:"user_id"`
	Email         string `json:"email" form:"email"`
	Password      string `json:"password" form:"password"`
	OperationType uint   `json:"operation_type" form:"operation_type"`
	jwt.RegisteredClaims
}

// ReleaseEmailToken 生成email token
func ReleaseEmailToken(userId, operationType uint, email, password string) (string, error) {
	claim := &EmailMsg{
		UserId:        userId,
		Email:         email,
		Password:      password,
		OperationType: operationType,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "127.0.0.1",
			Subject:   "validate email",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
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

// ParseEmailToken 解析email token
func ParseEmailToken(tokenString string) (*jwt.Token, *EmailMsg, error) {
	claims := &EmailMsg{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	return token, claims, err
}
