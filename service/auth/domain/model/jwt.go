package model

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type JWT struct {
	AccessToken      string `valid:"notnull"`
	IDToken          string `valid:"notnull"`
	ExpiresIn        int    `valid:"notnull"`
	RefreshExpiresIn int    `valid:"notnull"`
	RefreshToken     string `valid:"notnull"`
	TokenType        string `valid:"notnull"`
	NotBeforePolicy  int    `valid:"notnull"`
	SessionState     string `valid:"notnull"`
	Scope            string `valid:"notnull"`
}

func (t *JWT) isValid() error {

	_, err := govalidator.ValidateStruct(t)
	return err
}

func NewJWT(accessToken, idToken, refreshToken, tokenType, sessionState, scope string, expiresIn, refreshExpiresIn, notBeforePolicy int) (*JWT, error) {

	jwt := &JWT{
		AccessToken:      accessToken,
		IDToken:          idToken,
		ExpiresIn:        expiresIn,
		RefreshExpiresIn: refreshExpiresIn,
		RefreshToken:     refreshToken,
		TokenType:        tokenType,
		NotBeforePolicy:  notBeforePolicy,
		SessionState:     sessionState,
		Scope:            scope,
	}

	err := jwt.isValid()
	if err != nil {
		return nil, err
	}

	return jwt, nil
}
