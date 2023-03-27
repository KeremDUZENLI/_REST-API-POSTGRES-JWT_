package token

import (
	"postgre-project/common/env"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	FirstName string
	LastName  string
	Email     string
	UserType  string
	Uid       string
	jwt.StandardClaims
}

func GenerateToken(firstName string, lastName string, email string, userType string) (token string, refreshToken string, err error) {
	var expiresAt int64 = time.Now().Local().Add(time.Hour * time.Duration(24)).Unix()

	claims := &SignedDetails{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		UserType:  userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	if token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(env.SECRET_KEY)); err != nil {
		return
	}

	if refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).
		SignedString([]byte(env.SECRET_KEY)); err != nil {
		return
	}

	return
}
