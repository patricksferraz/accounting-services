package service

import (
	"context"

	"github.com/c4ut/accounting-services/service/common/pb"
	"github.com/c4ut/accounting-services/service/time-record/domain/model"
	"go.elastic.co/apm"
)

type AuthService struct {
	Service pb.AuthServiceClient
}

func (a *AuthService) Verify(ctx context.Context, accessToken string) (*model.EmployeeClaims, error) {
	span, ctx := apm.StartSpan(ctx, "Verify", "service")
	defer span.End()

	req := &pb.FindEmployeeClaimsByTokenRequest{
		AccessToken: accessToken,
	}

	employee, err := a.Service.FindEmployeeClaimsByToken(ctx, req)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		return nil, err
	}

	employeeClaims, err := model.NewEmployeeClaims(employee.Id, employee.Roles)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		return nil, err
	}

	return employeeClaims, nil
}

func NewAuthService(service pb.AuthServiceClient) *AuthService {
	return &AuthService{
		Service: service,
	}
}
