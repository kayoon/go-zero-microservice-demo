package common

import (
	"go-zero-microservice-demo/define"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

// Md5 加密
func Md5(s string) string  {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// GenerateToken 获取token
func GenerateToken(id int, name, groupid  string) (string, error)  {
	uc := define.UserClaim{
		Id:             id,
		Name:           name,
		GroupId:		groupid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyzeToken 解析token
func AnalyzeToken(token string) (*define.UserClaim, error)  {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid{
		return uc, errors.New("token is invalid")
	}
	return uc, nil
}