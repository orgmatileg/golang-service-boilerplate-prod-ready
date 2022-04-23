package jwtd

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

type ClaimData struct {
	UserID int    `json:"UserID,omitempty"`
	Role   string `json:"Role,omitempty"`
}

// Claim struct
type Claim struct {
	Data ClaimData `json:"Data"`
	jwt.StandardClaims
}

// GenerateToken ...
func GenerateToken(c Claim, secretString string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString([]byte(secretString))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken ...
func ParseToken(tokenString string, secretString string) (token *jwt.Token, err error) {

	token, err = jwt.Parse(tokenString, func(jt *jwt.Token) (interface{}, error) {

		// Untuk mencegah JWT Signing method NONE attack
		// Maka pastikan untuk memvalidasi juga Algoritma signing nya
		if _, ok := jt.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jt.Header["alg"])
		}
		return []byte(secretString), nil
	})

	return token, err
}

// IsValidToken validate JWT Token
func IsValidToken(tokenString string, secretString string) (bool, error) {
	token, err := ParseToken(tokenString, secretString)
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}

// ParseClaim func
func ParseClaim(tokenString string, secretString string) (*Claim, error) {

	claims := Claim{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretString), nil
	})
	if err != nil {
		return nil, err
	}

	return &claims, nil
}
