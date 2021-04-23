package repository

import (
	"context"

	"github.com/patricksferraz/accounting-services/service/time-record/domain/model"
)

type AuthRepositoryInterface interface {
	Verify(ctx context.Context, accessToken string) (*model.EmployeeClaims, error)
}
