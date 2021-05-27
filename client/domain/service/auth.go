package service

import (
	"context"

	authmodel "github.com/c4ut/accounting-services/service/auth/domain/model"
	"github.com/c4ut/accounting-services/service/common/pb"
)

type AuthService struct {
	Service pb.AuthServiceClient
}

func (a *AuthService) Login(ctx context.Context, username, password string) (*authmodel.JWT, error) {
	req := &pb.LoginRequest{
		Username: username,
		Password: password,
	}

	res, err := a.Service.Login(ctx, req)
	if err != nil {
		return nil, err
	}

	jwt := &authmodel.JWT{
		AccessToken:      res.AccessToken,
		IDToken:          res.IdToken,
		ExpiresIn:        int(res.ExpiresIn),
		RefreshExpiresIn: int(res.RefreshExpiresIn),
		RefreshToken:     res.RefreshToken,
		TokenType:        res.TokenType,
		NotBeforePolicy:  int(res.NotBeforePolicy),
		SessionState:     res.SessionState,
		Scope:            res.Scope,
	}

	return jwt, nil
}

func (a *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*authmodel.JWT, error) {
	req := &pb.RefreshTokenRequest{
		RefreshToken: refreshToken,
	}

	res, err := a.Service.RefreshToken(ctx, req)
	if err != nil {
		return nil, err
	}

	jwt := &authmodel.JWT{
		AccessToken:      res.AccessToken,
		IDToken:          res.IdToken,
		ExpiresIn:        int(res.ExpiresIn),
		RefreshExpiresIn: int(res.RefreshExpiresIn),
		RefreshToken:     res.RefreshToken,
		TokenType:        res.TokenType,
		NotBeforePolicy:  int(res.NotBeforePolicy),
		SessionState:     res.SessionState,
		Scope:            res.Scope,
	}

	return jwt, nil
}

func (a *AuthService) FindEmployeeClaimsByToken(ctx context.Context, accessToken string) (*authmodel.EmployeeClaims, error) {
	req := &pb.FindEmployeeClaimsByTokenRequest{
		AccessToken: accessToken,
	}

	res, err := a.Service.FindEmployeeClaimsByToken(ctx, req)
	if err != nil {
		return nil, err
	}

	employeeClaims, err := authmodel.NewEmployeeClaims(res.Id, res.Roles)
	if err != nil {
		return nil, err
	}

	return employeeClaims, nil
}

func NewAuthService(service pb.AuthServiceClient) *AuthService {
	return &AuthService{
		Service: service,
	}
}
