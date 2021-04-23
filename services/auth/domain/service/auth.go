package service

import (
	"context"

	"github.com/patricksferraz/accounting-services/services/auth/domain/model"
	"github.com/patricksferraz/accounting-services/services/auth/domain/repository"
)

type AuthService struct {
	AuthRepository repository.AuthRepositoryInterface
}

func (a *AuthService) Login(ctx context.Context, username, password string) (*model.JWT, error) {
	auth, err := model.NewAuth(username, password)
	if err != nil {
		return nil, err
	}

	jwt, err := a.AuthRepository.Login(ctx, auth)
	if err != nil {
		return nil, err
	}

	return jwt, nil
}

func (a *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*model.JWT, error) {
	jwt, err := a.AuthRepository.RefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	return jwt, nil
}

func (a *AuthService) FindEmployeeClaimsByToken(ctx context.Context, accessToken string) (*model.EmployeeClaims, error) {
	employee, err := a.AuthRepository.FindEmployeeClaimsByToken(ctx, accessToken)
	if err != nil {
		return nil, err
	}
	return employee, nil
}
