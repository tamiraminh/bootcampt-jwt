package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID   int64  `json:"userId`
	Username string `json:"username`
	jwt.StandardClaims
}

func generateJWT() (string, error) {
	secret := "secret"

	claims := Claims{
		UserID:   123,
		Username: "John Doe",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "evermos",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateJWT(tokenString string) (*Claims, error) {
	secret := "secret"

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("JWT validation failed: %v", err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("JWT is not valid")
}

func main() {


	jwtToken, err := generateJWT()
	if err != nil {
		fmt.Println(err)
	}


	fmt.Println("jwtToken:")
	fmt.Println(jwtToken)
	fmt.Println()

	claims , err := validateJWT(jwtToken)
	if err != nil {
		fmt.Println(err)
	}


	fmt.Printf("UserID: %d\n", claims.UserID)
	fmt.Printf("Username: %s\n", claims.Username)
}