package repository

import (
	"context"

	"github.com/mitchellh/mapstructure"
	"github.com/patricksferraz/accounting-services/service/auth/domain/model"
	"github.com/patricksferraz/accounting-services/service/auth/infrastructure/db"
)

type AuthRepositoryDb struct {
	Db *db.Keycloak
}

func (a *AuthRepositoryDb) Login(ctx context.Context, auth *model.Auth) (*model.JWT, error) {

	jwt, err := a.Db.Client.Login(ctx, a.Db.ClientID, a.Db.ClientSecret, a.Db.Realm, auth.Username, auth.Password)
	if err != nil {
		return nil, err
	}

	return &model.JWT{
		AccessToken:      jwt.AccessToken,
		IDToken:          jwt.IDToken,
		ExpiresIn:        jwt.ExpiresIn,
		RefreshExpiresIn: jwt.RefreshExpiresIn,
		RefreshToken:     jwt.RefreshToken,
		TokenType:        jwt.TokenType,
		NotBeforePolicy:  jwt.NotBeforePolicy,
		SessionState:     jwt.SessionState,
		Scope:            jwt.Scope,
	}, nil
}

func (a *AuthRepositoryDb) RefreshToken(ctx context.Context, refreshToken string) (*model.JWT, error) {

	jwt, err := a.Db.Client.RefreshToken(ctx, refreshToken, a.Db.ClientID, a.Db.ClientSecret, a.Db.Realm)
	if err != nil {
		return nil, err
	}

	return &model.JWT{
		AccessToken:      jwt.AccessToken,
		IDToken:          jwt.IDToken,
		ExpiresIn:        jwt.ExpiresIn,
		RefreshExpiresIn: jwt.RefreshExpiresIn,
		RefreshToken:     jwt.RefreshToken,
		TokenType:        jwt.TokenType,
		NotBeforePolicy:  jwt.NotBeforePolicy,
		SessionState:     jwt.SessionState,
		Scope:            jwt.Scope,
	}, nil
}

func (a *AuthRepositoryDb) FindEmployeeClaimsByToken(ctx context.Context, accessToken string) (*model.EmployeeClaims, error) {

	jwt, _, err := a.Db.Client.DecodeAccessToken(ctx, accessToken, a.Db.Realm, "account")
	if err != nil {
		return nil, err
	}

	employeeClaims := new(model.EmployeeClaims)
	mapstructure.Decode(jwt.Claims, employeeClaims)

	type ResourceAccess struct {
		ResourceAccess map[string]map[string][]string `mapstructure:"resource_access"`
	}

	ra := new(ResourceAccess)
	mapstructure.Decode(jwt.Claims, ra)

	roles := ra.ResourceAccess["checkpoint"]["roles"]
	employeeClaims.Roles = roles

	return employeeClaims, nil
}
