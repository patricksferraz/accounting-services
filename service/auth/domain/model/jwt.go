package model

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type JWT struct {
	AccessToken      string `json:"access_token" valid:"notnull"`
	IDToken          string `json:"id_token" valid:"notnull"`
	ExpiresIn        int    `json:"expires_in" valid:"notnull"`
	RefreshExpiresIn int    `json:"refresh_expires_in,omitempty" valid:"notnull"`
	RefreshToken     string `json:"refresh_token" valid:"notnull"`
	TokenType        string `json:"token_type" valid:"notnull"`
	NotBeforePolicy  int    `json:"not_before_policy" valid:"notnull"`
	SessionState     string `json:"session_state" valid:"notnull"`
	Scope            string `json:"scope" valid:"notnull"`
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
