// coding: utf-8
// @Author : lryself
// @Date : 2021/4/8 0:52
// @Software: GoLand

package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	UserId   string `json:"user_id"`
}

func genToken(claims *JWTClaims, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (claims *JWTClaims) MakeToken(TokenExpireTime int, secret []byte) (string, error) {
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(TokenExpireTime)).Unix()
	return genToken(claims, secret)
}

func VerifyToken(strToken string, secret []byte) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, err
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, err
	}
	return claims, nil
}

func RefreshToken(strToken string, secret []byte) (string, error) {
	claims, err := VerifyToken(strToken, secret)
	if err != nil {
		return "", err
	}
	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
	signedToken, err := genToken(claims, secret)
	if err != nil {
		return "", err
	}
	return signedToken, err
}
