package service

import (
	"context"

	"github.com/patricksferraz/accounting-services/service/time-record/domain/model"
	"github.com/patricksferraz/accounting-services/service/time-record/domain/repository"
)

type AuthService struct {
	AuthRepository repository.AuthRepositoryInterface
}

func (a *AuthService) Verify(ctx context.Context, accessToken string) (*model.EmployeeClaims, error) {
	employeeClaims, err := a.AuthRepository.Verify(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	return employeeClaims, nil
}

func NewAuthService(authRepository repository.AuthRepositoryInterface) *AuthService {
	return &AuthService{
		AuthRepository: authRepository,
	}
}
