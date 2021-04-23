package repository

import (
	"context"

	"github.com/patricksferraz/accounting-services/service/auth/domain/model"
)

type AuthRepositoryInterface interface {
	Login(ctx context.Context, auth *model.Auth) (*model.JWT, error)
	RefreshToken(ctx context.Context, refreshToken string) (*model.JWT, error)
	FindEmployeeClaimsByToken(ctx context.Context, accessToken string) (*model.EmployeeClaims, error)
}
