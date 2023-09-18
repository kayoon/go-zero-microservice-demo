package define

import (
	"github.com/dgrijalva/jwt-go"
)

// UserClaim 用户属性
type UserClaim struct {
	Id       int
	Identity string
	Name     string
	GroupId string
	jwt.StandardClaims
}

// JwtKey jwt密钥
var JwtKey = "demo-key"