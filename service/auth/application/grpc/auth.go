package grpc

import (
	"context"

	"github.com/patricksferraz/accounting-services/service/auth/domain/service"
	"github.com/patricksferraz/accounting-services/service/common/pb"
)

type AuthGrpcService struct {
	pb.UnimplementedAuthServiceServer
	AuthService *service.AuthService
}

func (a *AuthGrpcService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.JWT, error) {
	jwt, err := a.AuthService.Login(ctx, in.Username, in.Password)
	if err != nil {
		return &pb.JWT{}, err
	}

	return &pb.JWT{
		AccessToken:      jwt.AccessToken,
		IdToken:          jwt.IDToken,
		ExpiresIn:        int64(jwt.ExpiresIn),
		RefreshExpiresIn: int64(jwt.RefreshExpiresIn),
		RefreshToken:     jwt.RefreshToken,
		TokenType:        jwt.TokenType,
		NotBeforePolicy:  int64(jwt.NotBeforePolicy),
		SessionState:     jwt.SessionState,
		Scope:            jwt.Scope,
	}, err
}

func (a *AuthGrpcService) RefreshToken(ctx context.Context, in *pb.RefreshTokenRequest) (*pb.JWT, error) {
	jwt, err := a.AuthService.RefreshToken(ctx, in.RefreshToken)
	if err != nil {
		return &pb.JWT{}, err
	}

	return &pb.JWT{
		AccessToken:      jwt.AccessToken,
		IdToken:          jwt.IDToken,
		ExpiresIn:        int64(jwt.ExpiresIn),
		RefreshExpiresIn: int64(jwt.RefreshExpiresIn),
		RefreshToken:     jwt.RefreshToken,
		TokenType:        jwt.TokenType,
		NotBeforePolicy:  int64(jwt.NotBeforePolicy),
		SessionState:     jwt.SessionState,
		Scope:            jwt.Scope,
	}, err
}

func (a *AuthGrpcService) FindEmployeeClaimsByToken(ctx context.Context, in *pb.FindEmployeeClaimsByTokenRequest) (*pb.EmployeeClaims, error) {
	employee, err := a.AuthService.FindEmployeeClaimsByToken(ctx, in.AccessToken)
	if err != nil {
		return &pb.EmployeeClaims{}, err
	}

	return &pb.EmployeeClaims{
		Id:    employee.ID,
		Roles: employee.Roles,
	}, nil
}

func NewAuthGrpcService(service *service.AuthService) *AuthGrpcService {
	return &AuthGrpcService{
		AuthService: service,
	}
}
