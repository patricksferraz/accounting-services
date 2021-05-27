package repository

import (
	"context"

	"github.com/c4ut/accounting-services/service/auth/domain/model"
	"github.com/c4ut/accounting-services/service/auth/infrastructure/external"
	"github.com/mitchellh/mapstructure"
)

type AuthRepository struct {
	Service *external.Keycloak
}

func (a *AuthRepository) Login(ctx context.Context, auth *model.Auth) (*model.JWT, error) {

	jwt, err := a.Service.Client.Login(ctx, a.Service.ClientID, a.Service.ClientSecret, a.Service.Realm, auth.Username, auth.Password)
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

func (a *AuthRepository) RefreshToken(ctx context.Context, refreshToken string) (*model.JWT, error) {

	jwt, err := a.Service.Client.RefreshToken(ctx, refreshToken, a.Service.ClientID, a.Service.ClientSecret, a.Service.Realm)
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

func (a *AuthRepository) FindEmployeeClaimsByToken(ctx context.Context, accessToken string) (*model.EmployeeClaims, error) {

	jwt, _, err := a.Service.Client.DecodeAccessToken(ctx, accessToken, a.Service.Realm, a.Service.Audience)
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

	roles := ra.ResourceAccess[a.Service.ClientID]["roles"]
	employeeClaims.Roles = roles

	return employeeClaims, nil
}

func NewAuthRepository(service *external.Keycloak) *AuthRepository {
	return &AuthRepository{
		Service: service,
	}
}
