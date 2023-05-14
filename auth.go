package main

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var jwtKey = []byte("Cacique_Secret_2023")

type JWTClaim struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateJWT(password string, username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(8 * time.Hour)
	claims := &JWTClaim{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("No se puede canjear el secreto")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("Secreto expirado, inicie sesion para obtener un secreto nuevo")
		return
	}

	return

}
