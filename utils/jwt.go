package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtkey = []byte("dddjjd22dad")

type Claims struct {
	UserId int `json:"userId,omitempty"`
	jwt.StandardClaims
}

const tokenTTL = 24 * time.Hour

func GetToken(id int) (string, error) {
	expireTime := time.Now().Add(tokenTTL)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserId: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})
	str, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println("Get Token Err:", err)
		return "", err
	}
	return str, nil
}

func ParseToken(str string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(str, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})
	return token, claims, err
}
