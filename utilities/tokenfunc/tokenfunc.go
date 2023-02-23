package tokenfunc

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Id        uint   `json:"id"`
	Username  string `json:"userame"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

// 根据用户信息签发token
func GenerateToken(c Claims, secret string) (string, error) {
	now := time.Now()
	expireTime := now.Add(24 * time.Hour)
	claim := Claims{
		Id:        c.Id,
		Username:  c.Username,
		Authority: c.Authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "task",
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(secret))
	return token, err
}

// 解析token
func ParseToken(tokenString string, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !(ok && token.Valid) {
		return nil, err
	}
	return claims, nil
}
