package token

import (
	"postgre-project/common/env"
	"postgre-project/database"
	"postgre-project/database/model"
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

func GenerateToken(firstName string, lastName string, email string, userType string) (token string, err error) {
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

	if token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(env.SECRET_KEY)); err != nil {
		return
	}

	return
}

func ValidateToken(signedToken string) (*SignedDetails, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		keyFunction,
	)
	if err != nil {
		return nil, err
	}

	return token.Claims.(*SignedDetails), nil
}

func UpdateToken(id uint, signedToken string) error {
	updatedAt := time.Now()

	return database.ConnectDB().
		Table(model.TABLE).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"token":      signedToken,
			"updated_at": updatedAt,
		}).Error
}

func keyFunction(token *jwt.Token) (any, error) {
	return []byte(env.SECRET_KEY), nil
}
