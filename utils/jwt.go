package utils

import (
	"fmt"
	"time"
	. "websiteGin/model"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("dddjjd22dad")

type Claims struct {
	ID    int    `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	jwt.StandardClaims
}

const tokenTTL = 24 * time.Hour

func GetToken(u User) (string, error) {
	expireTime := time.Now().Add(tokenTTL)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		ID:    u.ID,
		Email: u.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})
	str, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println("Get Token Err:", err)
		return "", err
	}
	return str, nil
}

func ParseToken(str string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(str, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
