package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("FanOne")

type Claims struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

func GenerateToken(id int64, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		ID:       id,
		Username: username,
		//Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "MyTikTok",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
