package util

import (
	"gin_example/src/gin-blog/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret + strconv.Itoa(GenRand())[1:6])

type Claims struct {
	Username string `json:"username"`
	PubDesc  int    `json:"pub_desc"`
	jwt.StandardClaims
}

func GenerateToken(username string, pubDesc int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		Username: username,
		PubDesc:  pubDesc,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
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
