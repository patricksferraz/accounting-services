package grpc

import (
	"context"
	"log"

	"github.com/c4ut/accounting-services/service/time-record/domain/model"
	"github.com/c4ut/accounting-services/service/time-record/domain/service"
	"go.elastic.co/apm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	AuthService    *service.AuthService
	EmployeeClaims *model.EmployeeClaims
}

func (a *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		span, ctx := apm.StartSpan(ctx, "Unary", "application")
		defer span.End()

		log.Println("--> unary interceptor: ", info.FullMethod)

		err := a.authorize(ctx, info.FullMethod)
		if err != nil {
			apm.CaptureError(ctx, err).Send()
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (a *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		span, ctx := apm.StartSpan(ss.Context(), "Stream", "application")
		defer span.End()

		log.Println("--> stream interceptor: ", info.FullMethod)

		err := a.authorize(ctx, info.FullMethod)
		if err != nil {
			apm.CaptureError(ctx, err).Send()
			return err
		}

		return handler(srv, ss)
	}
}

func (a *AuthInterceptor) authorize(ctx context.Context, method string) error {
	span, ctx := apm.StartSpan(ctx, "authorize", "application")
	defer span.End()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err := status.Error(codes.Unauthenticated, "metadata is not provided")
		apm.CaptureError(ctx, err).Send()
		return err
	}

	values := md["authorization"]
	if len(values) == 0 {
		err := status.Error(codes.Unauthenticated, "authorization token is not provided")
		apm.CaptureError(ctx, err).Send()
		return err
	}

	accessToken := values[0]
	claims, err := a.AuthService.Verify(ctx, accessToken)
	if err != nil {
		err := status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
		apm.CaptureError(ctx, err).Send()
		return err
	}

	a.EmployeeClaims = claims

	for _, role := range claims.Roles {
		if role == method {
			return nil
		}
	}

	err = status.Error(codes.PermissionDenied, "no permission to access this RPC")
	apm.CaptureError(ctx, err).Send()
	return err
}

func NewAuthInterceptor(authService *service.AuthService) *AuthInterceptor {
	return &AuthInterceptor{
		AuthService: authService,
	}
}
