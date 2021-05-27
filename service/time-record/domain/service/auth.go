package service

import (
	"context"

	"github.com/c4ut/accounting-services/service/common/pb"
	"github.com/c4ut/accounting-services/service/time-record/domain/model"
)

type AuthService struct {
	Service pb.AuthServiceClient
}

func (a *AuthService) Verify(ctx context.Context, accessToken string) (*model.EmployeeClaims, error) {
	req := &pb.FindEmployeeClaimsByTokenRequest{
		AccessToken: accessToken,
	}

	employee, err := a.Service.FindEmployeeClaimsByToken(ctx, req)
	if err != nil {
		return nil, err
	}

	employeeClaims, err := model.NewEmployeeClaims(employee.Id, employee.Roles)
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
