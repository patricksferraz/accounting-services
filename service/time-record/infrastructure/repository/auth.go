package repository

import (
	"context"

	"github.com/patricksferraz/accounting-services/service/common/pb"
	"github.com/patricksferraz/accounting-services/service/time-record/domain/model"
)

type AuthRepositoryDb struct {
	Db pb.AuthServiceClient
}

func (a *AuthRepositoryDb) Verify(ctx context.Context, accessToken string) (*model.EmployeeClaims, error) {

	req := &pb.FindEmployeeClaimsByTokenRequest{
		AccessToken: accessToken,
	}

	employee, err := a.Db.FindEmployeeClaimsByToken(ctx, req)
	if err != nil {
		return nil, err
	}

	employeeClaims, err := model.NewEmployeeClaims(employee.Id, employee.Roles)
	if err != nil {
		return nil, err
	}

	return employeeClaims, nil
}

func NewAuthRepositoryDb(database pb.AuthServiceClient) *AuthRepositoryDb {
	return &AuthRepositoryDb{
		Db: database,
	}
}
